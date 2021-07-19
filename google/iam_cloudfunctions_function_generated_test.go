// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudFunctionsCloudFunctionIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"zip_path":      createZIPArchiveForCloudFunctionSource(t, testHTTPTriggerPath),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudFunctionsCloudFunctionIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloudfunctions_function_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-function%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccCloudFunctionsCloudFunctionIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_cloudfunctions_function_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-function%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudFunctionsCloudFunctionIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"zip_path":      createZIPArchiveForCloudFunctionSource(t, testHTTPTriggerPath),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccCloudFunctionsCloudFunctionIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloudfunctions_function_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s roles/viewer user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-function%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudFunctionsCloudFunctionIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"zip_path":      createZIPArchiveForCloudFunctionSource(t, testHTTPTriggerPath),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudFunctionsCloudFunctionIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloudfunctions_function_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-function%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudFunctionsCloudFunctionIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_cloudfunctions_function_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-function%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudFunctionsCloudFunctionIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name = "tf-test-cloudfunctions-function-example-bucket%{random_suffix}"
}

resource "google_storage_bucket_object" "archive" {
  name   = "index.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"
}

resource "google_cloudfunctions_function" "function" {
  name        = "tf-test-my-function%{random_suffix}"
  description = "My function"
  runtime     = "nodejs10"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.archive.name
  trigger_http          = "SECURE_ALWAYS"
  timeout               = 60
  entry_point           = "helloGET"
}

resource "google_cloudfunctions_function_iam_member" "foo" {
  project = google_cloudfunctions_function.function.project
  region = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccCloudFunctionsCloudFunctionIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name = "tf-test-cloudfunctions-function-example-bucket%{random_suffix}"
}

resource "google_storage_bucket_object" "archive" {
  name   = "index.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"
}

resource "google_cloudfunctions_function" "function" {
  name        = "tf-test-my-function%{random_suffix}"
  description = "My function"
  runtime     = "nodejs10"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.archive.name
  trigger_http          = "SECURE_ALWAYS"
  timeout               = 60
  entry_point           = "helloGET"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_cloudfunctions_function_iam_policy" "foo" {
  project = google_cloudfunctions_function.function.project
  region = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudFunctionsCloudFunctionIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name = "tf-test-cloudfunctions-function-example-bucket%{random_suffix}"
}

resource "google_storage_bucket_object" "archive" {
  name   = "index.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"
}

resource "google_cloudfunctions_function" "function" {
  name        = "tf-test-my-function%{random_suffix}"
  description = "My function"
  runtime     = "nodejs10"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.archive.name
  trigger_http          = "SECURE_ALWAYS"
  timeout               = 60
  entry_point           = "helloGET"
}

data "google_iam_policy" "foo" {
}

resource "google_cloudfunctions_function_iam_policy" "foo" {
  project = google_cloudfunctions_function.function.project
  region = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudFunctionsCloudFunctionIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name = "tf-test-cloudfunctions-function-example-bucket%{random_suffix}"
}

resource "google_storage_bucket_object" "archive" {
  name   = "index.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"
}

resource "google_cloudfunctions_function" "function" {
  name        = "tf-test-my-function%{random_suffix}"
  description = "My function"
  runtime     = "nodejs10"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.archive.name
  trigger_http          = "SECURE_ALWAYS"
  timeout               = 60
  entry_point           = "helloGET"
}

resource "google_cloudfunctions_function_iam_binding" "foo" {
  project = google_cloudfunctions_function.function.project
  region = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccCloudFunctionsCloudFunctionIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name = "tf-test-cloudfunctions-function-example-bucket%{random_suffix}"
}

resource "google_storage_bucket_object" "archive" {
  name   = "index.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"
}

resource "google_cloudfunctions_function" "function" {
  name        = "tf-test-my-function%{random_suffix}"
  description = "My function"
  runtime     = "nodejs10"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.archive.name
  trigger_http          = "SECURE_ALWAYS"
  timeout               = 60
  entry_point           = "helloGET"
}

resource "google_cloudfunctions_function_iam_binding" "foo" {
  project = google_cloudfunctions_function.function.project
  region = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
