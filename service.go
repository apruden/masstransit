package main

import (
	"database/sql"
	"fmt"
)

type RowScanner func(rows *sql.Rows) interface{}

func ExecuteQuery(db *sql.DB, sql_query string) *sql.Rows {
	rows, err := db.Query(sql_query)

	if err != nil {
		panic(err)
	}

	return rows
}

func ReadData(db *sql.DB, sql_query string, scanner RowScanner) []interface{} {
	rows := ExecuteQuery(db, sql_query)
	defer rows.Close()

	var result []interface{}

	for rows.Next() {
		item := scanner(rows)
		result = append(result, item)
	}

	return result
}

func ReadStop(db *sql.DB) []Stop {
	sql_query := `
	SELECT Id, Name, Code, Lat, Lon, Routes FROM Stops
	`
	var res []Stop

	items := ReadData(db, sql_query, func(rows *sql.Rows) interface{} {
		item := Stop{}
		err2 := rows.Scan(&item.Id, &item.Name, &item.Code, &item.Lat, &item.Lon, &item.Routes)

		if err2 != nil {
			panic(err2)
		}

		return item
	})

	for _, item := range items {
		res = append(res, item.(Stop))
	}

	return res
}

func ReadRouteShape(db *sql.DB, route string) []RouteShape {
	sql_query := fmt.Sprintf("SELECT Route, Points FROM RouteShapes WHERE Route = '%s'", route)
	var res []RouteShape

	items := ReadData(db, sql_query, func(rows *sql.Rows) interface{} {
		item := RouteShape{}
		err2 := rows.Scan(&item.Route, &item.Points)

		if err2 != nil {
			panic(err2)
		}

		return item
	})

	for _, item := range items {
		res = append(res, item.(RouteShape))
	}

	return res
}

func ReadRoute(db *sql.DB) []Route {
	sql_query := `
	SELECT Id, Code, Shape, Name, Stops FROM Routes
	`
	var res []Route

	items := ReadData(db, sql_query, func(rows *sql.Rows) interface{} {
		item := Route{}
		err2 := rows.Scan(&item.Id, &item.Code, &item.Shape, &item.Name, &item.Stops)

		if err2 != nil {
			panic(err2)
		}

		return item
	})

	for _, item := range items {
		res = append(res, item.(Route))
	}

	return res
}

func ReadSchedule(db *sql.DB, stopId string) []Schedule {
	sql_query := fmt.Sprintf("SELECT Stop, Route, Service, Code, Times FROM Schedules WHERE Stop = '%s'", stopId)

	var res []Schedule

	items := ReadData(db, sql_query, func(rows *sql.Rows) interface{} {
		item := Schedule{}
		err2 := rows.Scan(&item.Stop, &item.Route, &item.Service, &item.Code, &item.Times)

		if err2 != nil {
			panic(err2)
		}

		return item
	})

	for _, item := range items {
		res = append(res, item.(Schedule))
	}

	return res
}
