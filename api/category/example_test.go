// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package category_test

import (
	"fmt"

	"go.bmvs.io/ynab/api/category"

	"reflect"

	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api"
)

func ExampleService_GetCategory() {
	client := ynab.NewClient("<valid_ynab_access_token>")
	c, _ := client.Category().GetCategory("<valid_budget_id>", "<valid_category_id>")
	fmt.Println(reflect.TypeOf(c))

	// Output: *category.Category
}

func ExampleService_GetCategories() {
	client := ynab.NewClient("<valid_ynab_access_token>")
	categories, _ := client.Category().GetCategories("<valid_budget_id>")
	fmt.Println(reflect.TypeOf(categories))

	// Output: []*category.GroupWithCategories
}

func ExampleService_GetCategoryForMonth() {
	client := ynab.NewClient("<valid_ynab_access_token>")
	c, _ := client.Category().GetCategoryForMonth("<valid_budget_id>",
		"<valid_category_id>", api.Date{})
	fmt.Println(reflect.TypeOf(c))

	// Output: *category.Category
}

func ExampleService_GetCategoryForCurrentMonth() {
	client := ynab.NewClient("<valid_ynab_access_token>")
	c, _ := client.Category().GetCategoryForCurrentMonth("<valid_budget_id>",
		"<valid_category_id>")
	fmt.Println(reflect.TypeOf(c))

	// Output: *category.Category
}

func ExampleService_UpdateCategoryForMonth() {
	validMonth, _ := api.DateFromString("2018-01-01")
	validPayload := category.PayloadMonthCategory{Budgeted: 1000}

	client := ynab.NewClient("<valid_ynab_access_token>")
	c, _ := client.Category().UpdateCategoryForMonth("<valid_budget_id>",
		"<valid_category_id>", validMonth, validPayload)
	fmt.Println(reflect.TypeOf(c))

	// Output: *category.Category
}

func ExampleService_UpdateCategoryForCurrentMonth() {
	validPayload := category.PayloadMonthCategory{Budgeted: 1000}

	client := ynab.NewClient("<valid_ynab_access_token>")
	c, _ := client.Category().UpdateCategoryForCurrentMonth("<valid_budget_id>",
		"<valid_category_id>", validPayload)
	fmt.Println(reflect.TypeOf(c))

	// Output: *category.Category
}
