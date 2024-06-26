package repository

import (
	"context"
	"follower-service/model"
	"log"

	//"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type UserRepository struct {
	driver       neo4j.DriverWithContext
	logger       *log.Logger
	databaseName string
}

func NewUserRepository(logger *log.Logger) (*UserRepository, error) {
	uri := "bolt://follower-database:7687"
	user := "neo4j"
	pass := "password"
	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}

	return &UserRepository{
		driver:       driver,
		logger:       logger,
		databaseName: "neo4j",
	}, nil
}

func (ur *UserRepository) Follow(user, follower *model.User) error {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: ur.databaseName})
	defer session.Close(ctx)
	println(user.UserId)
	savedUser, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MERGE (u:User {userId: $userId}) "+
					"MERGE (f:User {userId: $followerId}) "+
					"CREATE (f)-[:FOLLOWS]->(u)",
				map[string]any{"userId": user.UserId, "followerId": follower.UserId})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		ur.logger.Println("Error inserting user:", err)
		return err
	}
	ur.logger.Println(savedUser)
	return nil
}

func (ur *UserRepository) Unfollow(user, unfollower *model.User) error {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: ur.databaseName})
	defer session.Close(ctx)

	savedUser, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (f:User {userId: $unfollowerId})-[r:FOLLOWS]->(u:User {userId: $userId}) "+
					"DELETE r",
				map[string]any{"userId": user.UserId, "unfollowerId": unfollower.UserId})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		ur.logger.Println("Error inserting user:", err)
		return err
	}
	ur.logger.Println(savedUser)
	return nil
}

func (ur *UserRepository) GetFollowers(user *model.User) ([]*model.User, error) {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// ExecuteRead for read transactions (Read and queries)
	userResults, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				`MATCH (f:User)-[:FOLLOWS]->(u:User {userId: $userID}) RETURN f`,
				map[string]interface{}{"userID": user.UserId})
			if err != nil {
				return nil, err
			}

			var followers []*model.User
			for result.Next(ctx) {
				record := result.Record()

				//log.Println(record)

				rawUser, ok := record.Get("f")
				if !ok {
					log.Println("Error: user node is missing")
					continue
				}
				//extracting node
				userNode, ok := rawUser.(neo4j.Node)

				if !ok {
					return nil, nil
				}

				userIdStr, ok := userNode.Props["userId"]
				if !ok || userIdStr == nil {
					// Handle the nil or missing userId appropriately
					log.Println("Error: userId is missing or nil")
					continue
				}

				followers = append(followers, &model.User{
					UserId: int(userIdStr.(int64)),
				})

			}
			return followers, nil

		})
	if err != nil {
		ur.logger.Println("Error querying search:", err)
		return nil, err
	}
	return userResults.([]*model.User), nil
}

func (ur *UserRepository) GetFollowing(user *model.User) ([]*model.User, error) {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: ur.databaseName})
	defer session.Close(ctx)

	userResults, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				`MATCH (u:User {userId: $userID})-[:FOLLOWS]->(f:User) RETURN f`,
				map[string]interface{}{"userID": user.UserId})
			if err != nil {
				return nil, err
			}

			var followers []*model.User
			for result.Next(ctx) {
				record := result.Record()

				//log.Println(record)

				rawUser, ok := record.Get("f")
				if !ok {
					log.Println("Error: user node is missing")
					continue
				}
				//extracting node
				userNode, ok := rawUser.(neo4j.Node)

				if !ok {
					return nil, nil
				}

				userIdStr, ok := userNode.Props["userId"]
				if !ok || userIdStr == nil {
					// Handle the nil or missing userId appropriately
					log.Println("Error: userId is missing or nil")
					continue
				}

				followers = append(followers, &model.User{
					UserId: int(userIdStr.(int64)),
				})

			}
			return followers, nil

		})

	if err != nil {
		ur.logger.Println("Error querying search:", err)
		return nil, err
	}
	return userResults.([]*model.User), nil

}

func (ur *UserRepository) GetFollowerSuggestions(user *model.User) ([]*model.User, error) {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: ur.databaseName})
	defer session.Close(ctx)

	userResults, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				`MATCH (u:User {userId: $userID})-[:FOLLOWS]->(f:User)-[:FOLLOWS]->(followed:User)
				 WHERE followed.userId <> u.userId AND NOT (u)-[:FOLLOWS]->(followed)
				 WITH followed, COUNT(f) AS numFollowers
				 RETURN followed
				 ORDER BY numFollowers DESC`,
				map[string]interface{}{"userID": user.UserId})
			if err != nil {
				return nil, err
			}

			var followers []*model.User
			for result.Next(ctx) {
				record := result.Record()

				//log.Println(record)

				rawUser, ok := record.Get("followed")
				if !ok {
					log.Println("Error: user node is missing")
					continue
				}
				//extracting node
				userNode, ok := rawUser.(neo4j.Node)

				if !ok {
					return nil, nil
				}

				userIdStr, ok := userNode.Props["userId"]
				if !ok || userIdStr == nil {
					// Handle the nil or missing userId appropriately
					log.Println("Error: userId is missing or nil")
					continue
				}

				followers = append(followers, &model.User{
					UserId: int(userIdStr.(int64)),
				})

			}
			return followers, nil

		})

	if err != nil {
		ur.logger.Println("Error querying search:", err)
		return nil, err
	}
	return userResults.([]*model.User), nil

}
