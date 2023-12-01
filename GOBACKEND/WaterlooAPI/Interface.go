package WaterlooAPI

import "encoding/json"

// Academic Organization

func GetAllAcademicOrganization() ([]byte, error) {
	Address := "/v3/AcademicOrganizations"
	return AbstractedGET(Address)
}

func GetSpecificAcademicOrganization(organizationCode string) ([]byte, error) {
	Address := "/v3/AcademicOrganizations" + organizationCode
	return AbstractedGET(Address)
}

// Account

/*
* Use this method to request an API key and begin your registration process
 */
func PostRegister(email, project, uri string) ([]byte, error) {
	requestBody, err := json.Marshal(struct {
		Email   string `json:"email"`
		Project string `json:"project"`
		URI     string `json:"uri"`
	}{
		Email:   email,
		Project: project,
		URI:     uri,
	})
	Address := "/v3/Account/Register"
	return AbstractedPOST(requestBody, err, Address)
}

/*
* Use this method to see if an email has already been registered for an API key
 */
func CheckEmailRegistration(email string) ([]byte, error) {
	Address := "/v3/Account/" + email
	return AbstractedGET(Address)
}

/*
* Use this method to have us re-send the confirmation email to an account pending confirmation, if it exists. The activation code will be reset in the process.
 */
func ResendConfirmation(email string) ([]byte, error) {
	Address := "/v3/Account/" + email + "/notify"
	return AbstractedGET(Address)
}

/*
* User this method to put your account in pending confirmation status and generate a new API key. Your old key will no longer grant access. The account will need to be confirmed again before the new key grants access.
 */
func GetNewAPIKey(email string) ([]byte, error) {
	Address := "/v3/Account/" + email + "/" + apiKey + "/reset"
	return AbstractedGET(Address)
}

/*
* Use this method to confirm your email and activate your account
* For the API KEY, ONLY USE AFTER POSTREGISTER()
 */
func ConfirmEmail(email, code string) ([]byte, error) {
	requestBody, err := json.Marshal(struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}{
		Email: email,
		Code:  code,
	})
	Address := "/v3/Account/Confirm"
	return AbstractedPOST(requestBody, err, Address)
}

// ClassSchedules

// Get the Course IDs that have one or more active and associated schedules in the given Term
func GetCurseID(termCode string) ([]byte, error) {
	Address := "/v3/ClassSchedules/" + termCode
	return AbstractedGET(Address)
}

// Get Class data for a scheduled Course by Course ID, in a specific Term
func GetClassDataByCourseID(termCode string, courseId string) ([]byte, error) {
	Address := "/v3/ClassSchedules/" + termCode + "/" + courseId
	return AbstractedGET(Address)
}

// Get Class data for a scheduled Course by Subject and catalog number, in a specific Term
func GetClassDataBySubjectAndCatalog(termCode string, subject string, catalogNumber string) ([]byte, error) {
	Address := "/v3/ClassSchedules/" + termCode + "/" + subject + "/" + catalogNumber
	return AbstractedGET(Address)
}

// Courses

// Get all Course data for a Term
func GetAllCourseData(termCode string) ([]byte, error) {
	Address := "/v3/Courses/" + termCode
	return AbstractedGET(Address)
}

// Get Course catalog data filtered by Term and Course ID, multiple Courses are usually cross listed
func GetCourseCatalogByCourseID(termCode string, courseId string) ([]byte, error) {
	Address := "/v3/Courses/" + termCode + "/" + courseId
	return AbstractedGET(Address)
}

// Get Course catalog data filtered by Term, Course ID, and offer number
func GetCourseCatalogByOfferNumber(termCode string, courseId string, offerNumber string) ([]byte, error) {
	Address := "/v3/Courses/" + termCode + "/" + courseId + "/" + offerNumber
	return AbstractedGET(Address)
}

// Get Course catalog data filtered by Term and Subject code
func GetCourseCatalogBySubject(termCode string, subject string) ([]byte, error) {
	Address := "/v3/Courses/" + termCode + "/" + subject
	return AbstractedGET(Address)
}

// Get Course catalog data filtered by Term, Subject, and catalog number
func GetCourseCatalogBySubjectAndCatalog(termCode string, subject string, catalogNumber string) ([]byte, error) {
	Address := "/v3/Courses/" + termCode + "/" + subject + "/" + catalogNumber
	return AbstractedGET(Address)
}

// ExamSchedules

// Returns Exam Schedule data for the current Term
func GetCurrentTermExamSchedule() ([]byte, error) {
	Address := "/v3/ExamSchedules"
	return AbstractedGET(Address)
}

// Returns Exam Schedule data for the requested Term
func GetExamScheduleByTerm(code string) ([]byte, error) {
	Address := "/v3/ExamSchedules/" + code
	return AbstractedGET(Address)
}

// Food Services

// Get all food service Outlet data
func GetAllFoodServiceOutlets() ([]byte, error) {
	Address := "/v3/FoodServices/outlets"
	return AbstractedGET(Address)
}

// Get specific food services Outlet data by Id
func GetFoodServiceOutletById(id string) ([]byte, error) {
	Address := "/v3/FoodServices/outlets/" + id
	return AbstractedGET(Address)
}

// Get specific food services Outlet data by Outlet name
func GetFoodServiceOutletByName(name string) ([]byte, error) {
	Address := "/v3/FoodServices/outlets/" + name
	return AbstractedGET(Address)
}

// Get all food service Franchise data
func GetAllFoodServiceFranchises() ([]byte, error) {
	Address := "/v3/FoodServices/franchises"
	return AbstractedGET(Address)
}

// Get specific food services Franchise data by Id
func GetFoodServiceFranchiseById(id string) ([]byte, error) {
	Address := "/v3/FoodServices/franchises/" + id
	return AbstractedGET(Address)
}

// Get specific food services Franchise data by Franchise name
func GetFoodServiceFranchiseByName(name string) ([]byte, error) {
	Address := "/v3/FoodServices/franchises/" + name
	return AbstractedGET(Address)
}

// HolidayDates

// Retrieves data for all paid holidays as published by Human Resources
func GetAllPaidHolidays() ([]byte, error) {
	Address := "/v3/HolidayDates/paidholidays"
	return AbstractedGET(Address)
}

// Retrieves data for paid holidays associated to the given year as published by Human Resources
func GetPaidHolidaysByYear(year string) ([]byte, error) {
	Address := "/v3/HolidayDates/paidholidays/" + year
	return AbstractedGET(Address)
}

// Retrieves data for all paid holidays as published by Human Resources, as an ICS format feed. Allows anonymous access.
func GetAllPaidHolidaysICS() ([]byte, error) {
	Address := "/v3/HolidayDates/paidholidays/ics"
	return AbstractedGET(Address)
}

// Important Dates

// Returns all current data for Important Dates
func GetAllImportantDates() ([]byte, error) {
	Address := "/v3/ImportantDates"
	return AbstractedGET(Address)
}

// Locations

// Get all building location data
func GetAllBuildingLocations() ([]byte, error) {
	Address := "/v3/Locations"
	return AbstractedGET(Address)
}

// Get all building location data as GEO JSON
func GetAllBuildingLocationsGeoJSON() ([]byte, error) {
	Address := "/v3/Locations/geojson"
	return AbstractedGET(Address)
}

// Gets building by matched building code, case insensitive
func GetBuildingByCode(locationCode string) ([]byte, error) {
	Address := "/v3/Locations/" + locationCode
	return AbstractedGET(Address)
}

// Gets building by matched building code, case insensitive, as GEO JSON
func GetBuildingByCodeGeoJSON(locationCode string) ([]byte, error) {
	Address := "/v3/Locations/" + locationCode + "/geojson"
	return AbstractedGET(Address)
}

// Gets buildings by matched building name, contains, case insensitive
func SearchBuildingByName(locationName string) ([]byte, error) {
	Address := "/v3/Locations/search/" + locationName
	return AbstractedGET(Address)
}

// Gets buildings by matched building name, contains, case insensitive, as GEO JSON
func SearchBuildingByNameGeoJSON(locationName string) ([]byte, error) {
	Address := "/v3/Locations/search/" + locationName + "/geojson"
	return AbstractedGET(Address)
}

// Subjects

// Gets all Subject data
func GetAllSubjects() ([]byte, error) {
	Address := "/v3/Subjects"
	return AbstractedGET(Address)
}

// Gets Subject data filtered by Subject code
func GetSubjectByCode(code string) ([]byte, error) {
	Address := "/v3/Subjects/" + code
	return AbstractedGET(Address)
}

// Gets Subject data for Subjects associated to an Academic Organization by Organization code
func GetSubjectsAssociatedToOrganization(organizationCode string) ([]byte, error) {
	Address := "/v3/Subjects/associatedto/" + organizationCode
	return AbstractedGET(Address)
}

// Terms

// Gets all Term data that is effective at the current time
func GetAllTerms() ([]byte, error) {
	Address := "/v3/Terms"
	return AbstractedGET(Address)
}

// Gets the current Term data
func GetCurrentTerm() ([]byte, error) {
	Address := "/v3/Terms/current"
	return AbstractedGET(Address)
}

// Gets Term data for a specific Term filtered by code
func GetTermByCode(code string) ([]byte, error) {
	Address := "/v3/Terms/" + code
	return AbstractedGET(Address)
}

// Gets Term data for terms that are part of a specific academic year
func GetTermsForAcademicYear(year string) ([]byte, error) {
	Address := "/v3/Terms/foracademicyear/" + year
	return AbstractedGET(Address)
}

// Wcms

// Retrieves information about all active and published WCMS sites
func GetAllWCMSInformation() ([]byte, error) {
	Address := "/v3/Wcms"
	return AbstractedGET(Address)
}

// Retrieves information about a specific WCMS site by Site Id
func GetWCMSById(id string) ([]byte, error) {
	Address := "/v3/Wcms/" + id
	return AbstractedGET(Address)
}

// Retrieves the latest news items across all WCMS sites, ordered by posted date
func GetLatestNewsItems(maxItems string) ([]byte, error) {
	Address := "/v3/Wcms/latestnews/" + maxItems
	return AbstractedGET(Address)
}

// Retrieves the latest events items across all WCMS sites, ordered by event start date
func GetLatestEventItems(maxItems string) ([]byte, error) {
	Address := "/v3/Wcms/latestevents/" + maxItems
	return AbstractedGET(Address)
}

// Retrieves the latest blog post items across all WCMS sites, ordered by posted date
func GetLatestBlogPostItems(maxItems string) ([]byte, error) {
	Address := "/v3/Wcms/latestposts/" + maxItems
	return AbstractedGET(Address)
}

// Retrieves the latest opportunity items across all WCMS sites, ordered by posted/open date
func GetLatestOpportunityItems(maxItems string) ([]byte, error) {
	Address := "/v3/Wcms/latestopportunities/" + maxItems
	return AbstractedGET(Address)
}

// Retrieves all news items for a specific WCMS site by Site Id
func GetNewsItemsForWCMSById(id string) ([]byte, error) {
	Address := "/v3/Wcms/" + id + "/news"
	return AbstractedGET(Address)
}

// Retrieves all blog post items for a specific WCMS site by Site Id
func GetBlogPostsForWCMSById(id string) ([]byte, error) {
	Address := "/v3/Wcms/" + id + "/posts"
	return AbstractedGET(Address)
}

// Retrieves all opportunity items for a specific WCMS site by Site
func GetOpportunities(siteId string) ([]byte, error) {
	address := "/v3/Wcms/" + siteId + "/opportunities"
	return AbstractedGET(address)
}

// retrieves all event items for a specific WCMS site by Site Id.
func GetEvents(siteId string) ([]byte, error) {
	address := "/v3/Wcms/" + siteId + "/events"
	return AbstractedGET(address)
}
