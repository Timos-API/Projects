package projects

import (
	"Timos-API/Projects/database"
	ctx "context"
	"fmt"
	"strconv"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func collection() *mongo.Collection {
	return database.Database.Collection("projects")
}

func printError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(Exception{message})
}

func getProjectById(id string) *Project {
	var project Project
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Printf("Invalid ObjectID %v\n", id)
		return nil
	}

	err = collection().FindOne(ctx.Background(), bson.M{"_id": oid}).Decode(&project)

	if err != nil {
		fmt.Printf("Project not found... (%v) %v \n", id, err)
		return nil
	}

	return &project
}

func getAll(w http.ResponseWriter, req *http.Request, filter primitive.M) {
	options := options.Find()
	query := req.URL.Query()
	qLimit, qSkip, qQuery := query.Get("limit"), query.Get("skip"), query.Get("query")

	if len(qLimit) > 0 {
		if limit, err := strconv.ParseInt(qLimit, 10, 64); err == nil {
			options.SetLimit(limit)
		}
	}

	if len(qSkip) > 0 {
		if skip, err := strconv.ParseInt(qSkip, 10, 64); err == nil {
			options.SetSkip(skip)
		}
	}

	if len(qQuery) > 0 {
		regex := primitive.Regex{Pattern: qQuery, Options: "i"}
		filter = bson.M{
			"$and": []bson.M{
				filter,
				{"$or": []bson.M{
					{"title": regex},
					{"icon": regex},
					{"description": regex},
					{"hero": regex},
					{"thumbnail": regex},
					{"website": regex},
					{"github": regex},
					{"designTools": regex},
					{"development": regex},
					{"frameworks": regex},
				}},
			},
		}
	}

	cursor, err := collection().Find(ctx.Background(), filter, options)

	if err != nil {
		printError(w, err.Error())
		return
	}

	defer cursor.Close(ctx.Background())

	projects := []Project{}
	for cursor.Next(ctx.Background()) {
		var project Project
		cursor.Decode(&project)
		projects = append(projects, project)
	}

	if err := cursor.Err(); err != nil {
		printError(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(projects)
}

func getAllProjects(w http.ResponseWriter, req *http.Request) {
	getAll(w, req, bson.M{})
}

func getProject(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	project := getProjectById(params["id"])

	if project == nil {
		printError(w, "Project not found")
		return
	}

	json.NewEncoder(w).Encode(project)
}
