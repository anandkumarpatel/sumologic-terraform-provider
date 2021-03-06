// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated. Do not submit
//     changes to this file.
//
// ----------------------------------------------------------------------------\
package sumologic

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSumologicUser_basic(t *testing.T) {
	var user User
	testFirstName := FieldsMap["User"]["firstName"]
	testLastName := FieldsMap["User"]["lastName"]
	testEmail := FieldsMap["User"]["email"]
	testIsActive, _ := strconv.ParseBool(FieldsMap["User"]["isActive"])

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy(user),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSumologicUserConfigImported(testFirstName, testLastName, testEmail, testIsActive),
			},
			{
				ResourceName:      "sumologic_user.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccUser_create(t *testing.T) {
	var user User
	testFirstName := FieldsMap["User"]["firstName"]
	testLastName := FieldsMap["User"]["lastName"]
	testEmail := FieldsMap["User"]["email"]
	testIsActive, _ := strconv.ParseBool(FieldsMap["User"]["isActive"])
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy(user),
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicUser(testFirstName, testLastName, testEmail, testIsActive),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("sumologic_user.test", &user, t),
					testAccCheckUserAttributes("sumologic_user.test"),
					resource.TestCheckResourceAttr("sumologic_user.test", "first_name", testFirstName),
					resource.TestCheckResourceAttr("sumologic_user.test", "last_name", testLastName),
					resource.TestCheckResourceAttr("sumologic_user.test", "email", testEmail),
					resource.TestCheckResourceAttr("sumologic_user.test", "is_active", strconv.FormatBool(testIsActive)),
				),
			},
		},
	})
}

func testAccCheckUserDestroy(user User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Client)
		for _, r := range s.RootModule().Resources {
			id := r.Primary.ID
			u, err := client.GetUser(id)
			if err != nil {
				return fmt.Errorf("Encountered an error: " + err.Error())
			}
			if u != nil {
				return fmt.Errorf("User still exists")
			}
		}
		return nil
	}
}

func testAccCheckUserExists(name string, user *User, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			//need this so that we don't get an unused import error for strconv in some cases
			return fmt.Errorf("Error = %s. User not found: %s", strconv.FormatBool(ok), name)
		}

		//need this so that we don't get an unused import error for strings in some cases
		if strings.EqualFold(rs.Primary.ID, "") {
			return fmt.Errorf("User ID is not set")
		}

		id := rs.Primary.ID
		c := testAccProvider.Meta().(*Client)
		newUser, err := c.GetUser(id)
		if err != nil {
			return fmt.Errorf("User %s not found", id)
		}
		user = newUser
		return nil
	}
}

func TestAccUser_update(t *testing.T) {
	var user User
	testFirstName := FieldsMap["User"]["firstName"]
	testLastName := FieldsMap["User"]["lastName"]
	testEmail := FieldsMap["User"]["email"]
	testIsActive, _ := strconv.ParseBool(FieldsMap["User"]["isActive"])

	testUpdatedFirstName := FieldsMap["User"]["updatedFirstName"]
	testUpdatedLastName := FieldsMap["User"]["updatedLastName"]
	testUpdatedEmail := FieldsMap["User"]["updatedEmail"]
	testUpdatedIsActive, _ := strconv.ParseBool(FieldsMap["User"]["updatedIsActive"])

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy(user),
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicUser(testFirstName, testLastName, testEmail, testIsActive),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("sumologic_user.test", &user, t),
					testAccCheckUserAttributes("sumologic_user.test"),
					resource.TestCheckResourceAttr("sumologic_user.test", "first_name", testFirstName),
					resource.TestCheckResourceAttr("sumologic_user.test", "last_name", testLastName),
					resource.TestCheckResourceAttr("sumologic_user.test", "email", testEmail),
					resource.TestCheckResourceAttr("sumologic_user.test", "is_active", strconv.FormatBool(testIsActive)),
				),
			},
			{
				Config: testAccSumologicUserUpdate(testUpdatedFirstName, testUpdatedLastName, testUpdatedEmail, testUpdatedIsActive),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("sumologic_user.test", &user, t),
					testAccCheckUserAttributes("sumologic_user.test"),
					resource.TestCheckResourceAttr("sumologic_user.test", "first_name", testUpdatedFirstName),
					resource.TestCheckResourceAttr("sumologic_user.test", "last_name", testUpdatedLastName),
					resource.TestCheckResourceAttr("sumologic_user.test", "email", testUpdatedEmail),
					resource.TestCheckResourceAttr("sumologic_user.test", "is_active", strconv.FormatBool(testUpdatedIsActive)),
				),
			},
		},
	})
}

func testAccCheckSumologicUserConfigImported(firstName string, lastName string, email string, isActive bool) string {
	return fmt.Sprintf(`
resource "sumologic_role" "testRole" {
	name = "testRole Name"
	capabilities = []
	description = "testRole Description"
	filter_predicate = ""
}
resource "sumologic_user" "foo" {
      first_name = "%s"
      last_name = "%s"
      email = "%s"
      role_ids = ["${sumologic_role.testRole.id}"]
      is_active = %t
}
`, firstName, lastName, email, isActive)
}

func testAccSumologicUser(firstName string, lastName string, email string, isActive bool) string {
	return fmt.Sprintf(`
resource "sumologic_role" "testRole" {
	name = "testRole Name"
	capabilities = []
	description = "testRole Description"
	filter_predicate = ""
}
resource "sumologic_user" "test" {
    first_name = "%s"
    last_name = "%s"
    email = "%s"
    role_ids = ["${sumologic_role.testRole.id}"]
    is_active = %t
}
`, firstName, lastName, email, isActive)
}

func testAccSumologicUserUpdate(firstName string, lastName string, email string, isActive bool) string {
	return fmt.Sprintf(`
resource "sumologic_role" "testRole" {
	name = "testRole Name"
	capabilities = []
	description = "testRole Description"
	filter_predicate = ""
}
resource "sumologic_user" "test" {
      first_name = "%s"
      last_name = "%s"
      email = "%s"
      role_ids = ["${sumologic_role.testRole.id}"]
      is_active = %t
}
`, firstName, lastName, email, isActive)
}

func testAccCheckUserAttributes(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		f := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(name, "first_name"),
			resource.TestCheckResourceAttrSet(name, "last_name"),
			resource.TestCheckResourceAttrSet(name, "email"),
			resource.TestCheckResourceAttrSet(name, "is_active"),
		)
		return f(s)
	}
}
