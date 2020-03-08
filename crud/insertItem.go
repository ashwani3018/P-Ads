package crud

import (
	"Advertisement/constant"
	"Advertisement/model"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"strconv"
)

func InsertItem() {
	// snippet-start:[dynamodb.go.create_item.session]
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	// snippet-end:[dynamodb.go.create_item.session]

	//////////////////////////////////////////////////////////////////


	// Each Category Logic
	categories := []model.Category{}
	for j := 0; j<3; j++ {
		// Category Ads Array
		catAds := []model.AdsItem{}
		// Category Size
		categorySize := 5
		// Each Category All Ads Initializing
		for i := 0; i < categorySize; i++ {
			categoryAds := model.AdsItem{Position: i, AdsID: "AdsID "+strconv.Itoa(j)+strconv.Itoa(i), Type: constant.ADS_TYPE_300x250}
			catAds = append(catAds, categoryAds)
		}
		// Each Sub-Category Init
		category := model.Category{ID : "categoryId "+strconv.Itoa(j), CategoryAds:catAds}
		categories = append(categories, category)
	}




	// Sub-Category Ads Logic
	subCategories := []model.SubCategory{}
	for j := 0; j<3; j++ {
		// Category Ads Array
		subcatAds := []model.AdsItem{}
		// Category Size
		categorySize := 5
		// Each Category All Ads Initializing
		for i := 0; i < categorySize; i++ {
			subCategoryAds := model.AdsItem{Position: i, AdsID: "AdsID "+strconv.Itoa(j)+strconv.Itoa(i), Type: constant.ADS_TYPE_300x250}
			subcatAds = append(subcatAds, subCategoryAds)
		}

		// Each Sub-Category Init
		subCategory := model.SubCategory{ID : "subCategoryId "+strconv.Itoa(j), CategoryAds:subcatAds}
		subCategories = append(subCategories, subCategory)
	}


	// Ads Struct Initialisation
	/*adsConfig := model.AdsConfig{
		ID:  100,
	}*/
	ads := model.Ads{ID: 300}

	ads.Category = categories
	ads.SubCategory = subCategories

	//adsConfig.Ads = ads
	//////////////////////////////////////////////////////////////////

	// snippet-start:[dynamodb.go.create_item.assign_struct]
	/*itemItem := model.TestItem{
		Year:   2015,
		Title:  "The Big New Movie",
		Plot:   "Nothing happens at all.",
		Rating: 0.0,
		ID: 90,
	}*/
	av, err := dynamodbattribute.MarshalMap(ads)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(constant.TableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//year := strconv.Itoa(item.Year)

	//fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)
	fmt.Println("Successfully added '" + constant.TableName)
	// snippet-end:[dynamodb.go.create_item.call]
}
