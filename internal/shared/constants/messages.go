package constants

const (
	// General Constants
	TIMEOUT_DURATION = 5 // Timeout duration in seconds

	// Error Messages
	ERR_FAILED_TO_GET_PROJECT_CATEGORIES = "Failed to get project categories"
	ERR_FAILED_TO_GET_CATEGORIES         = "Failed to get categories"
	ERR_FAILED_TO_GET_CATEGORY           = "Failed to get category"
	ERR_FAILED_TO_GET_CATEGORY_TASKS     = "Failed to get category tasks"

	ERR_FAILED_TO_GET_PROJECT  = "Failed to get project"
	ERR_FAILED_TO_CONVERT_ID   = "failed to convert inserted ID to ObjectID"
	ERR_CATEGORY_NOT_FOUND     = "category not found"
	ERR_INVALID_OBJECT_ID      = "invalid ObjectID format"
	ERR_INSERT_CATEGORY_FAILED = "failed to insert category"
	ERR_UPDATE_CATEGORY_FAILED = "failed to update category"
	ERR_DELETE_CATEGORY_FAILED = "failed to delete category"
	ERR_INSERT_PROJECT_FAILED  = "failed to insert project"
	ERR_UPDATE_PROJECT_FAILED  = "failed to update project"
	ERR_DELETE_PROJECT_FAILED  = "failed to delete project"
	ERR_PROJECT_NOT_FOUND      = "project not found"
	ERR_ADD_CATEGORY           = "failed to add category"
	ERR_UPDATE_CATEGORY        = "failed to update category"
	ERR_DELETE_CATEGORY        = "failed to delete category"
	SUCCESS_DELETE_CATEGORY    = "Category deleted successfully with Id"
	SUCCESS_ADD_CATEGORY       = "Category created successfully with Id"
	SUCCESS_UPDATE_CATEGORY    = "Category updated successfully with Id"
	ERR_ADD_PROJECT            = "failed to add project"
	ERR_UPDATE_PROJECT         = "failed to update project"
	ERR_DELETE_PROJECT         = "failed to delete project"
	SUCCESS_DELETE_PROJECT     = "Project deleted successfully with Id"
	SUCCESS_ADD_PROJECT        = "Project created successfully with Id"
	SUCCESS_UPDATE_PROJECT     = "Project updated successfully with Id"
	ERR_MONGO_CONNECTION       = "failed to connect to MongoDB"

	// Task-related error and success messages
	ERR_TASK_NOT_FOUND  = "task not found"
	ERR_ADD_TASK        = "failed to add task"
	ERR_UPDATE_TASK     = "failed to update task"
	ERR_DELETE_TASK     = "failed to delete task"
	SUCCESS_ADD_TASK    = "Task created successfully with Id"
	SUCCESS_UPDATE_TASK = "Task updated successfully with Id"
	SUCCESS_DELETE_TASK = "Task deleted successfully with Id"

	// User-related error and success messages
	ERR_USER_NOT_FOUND      = "user not found"
	ERR_REGISTER_USER       = "failed to register user"
	ERR_INVALID_CREDENTIALS = "invalid email or password"
	SUCCESS_REGISTER_USER   = "User registered successfully with Id"
	SUCCESS_LOGIN_USER      = "User logged in successfully with Id"
	SUCCESS_LOGOUT_USER     = "User logged out successfully"

	// Category Fields
	CATEGORY_NAME_FIELD  = "name"
	CATEGORY_TASKS_FIELD = "tasks"

	// Project Fields
	PROJECT_NAME_FIELD       = "name"
	PROJECT_CATEGORIES_FIELD = "categories"

	ERR_INVALID_INPUT                      = "Invalid input"
	ERR_RETRIEVE_TASKS                     = "Failed to retrieve tasks"
	ERR_LOGOUT_FAILED                      = "Failed to logout"
	ERR_ENV_FILE_NOT_FOUND                 = "No .env file found"
	ERR_MONGODB_URI_NOT_FOUND              = "MONGODB_URI not found in environment"
	SUCCESS_MONGODB_CONNECTION_ESTABLISHED = "MongoDB's Database Connection Successfully Established"

	// Repository error messages
	ERR_INSERT_FAILED    = "insert failed"
	ERR_ENTITY_NOT_FOUND = "entity not found"
	ERR_UPDATE_FAILED    = "update failed"
	ERR_DELETE_FAILED    = "delete failed"
)
