# Go API client for Moosend.Wrappers.GOWrapper

TODO: Add a description

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./Moosend.Wrappers.GOWrapper"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.moosend.com/v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CampaignsApi* | [**ABTestCampaignSummary**](docs/CampaignsApi.md#abtestcampaignsummary) | **Get** /campaigns/{CampaignID}/view_ab_summary.{Format} | AB Test Campaign Summary
*CampaignsApi* | [**ActivityByLocation**](docs/CampaignsApi.md#activitybylocation) | **Get** /campaigns/{CampaignID}/stats/countries.{Format} | Activity By Location
*CampaignsApi* | [**CampaignSummary**](docs/CampaignsApi.md#campaignsummary) | **Get** /campaigns/{CampaignID}/view_summary.{Format} | Campaign Summary
*CampaignsApi* | [**CloningAnExistingCampaign**](docs/CampaignsApi.md#cloninganexistingcampaign) | **Post** /campaigns/{CampaignID}/clone.{Format} | Cloning An Existing Campaign
*CampaignsApi* | [**CreatingADraftCampaign**](docs/CampaignsApi.md#creatingadraftcampaign) | **Post** /campaigns/create.{Format} | Creating A Draft Campaign
*CampaignsApi* | [**DeletingACampaign**](docs/CampaignsApi.md#deletingacampaign) | **Delete** /campaigns/{CampaignID}/delete.{Format} | Deleting A Campaign
*CampaignsApi* | [**GetAllCampaigns**](docs/CampaignsApi.md#getallcampaigns) | **Get** /campaigns.{Format} | Get All Campaigns
*CampaignsApi* | [**GetCampaignStatisticsWithPagingFiltered**](docs/CampaignsApi.md#getcampaignstatisticswithpagingfiltered) | **Get** /campaigns/{CampaignID}/stats/{Type}.{Format} | Get Campaign Statistics With Paging &amp; Filtered
*CampaignsApi* | [**GetCampaignsByPage**](docs/CampaignsApi.md#getcampaignsbypage) | **Get** /campaigns/{Page}.{Format} | Get Campaigns By Page
*CampaignsApi* | [**GetCampaignsByPageAndPagesize**](docs/CampaignsApi.md#getcampaignsbypageandpagesize) | **Get** /campaigns/{Page}/{PageSize}.{Format} | Get Campaigns By Page And Pagesize
*CampaignsApi* | [**GettingAllYourSenders**](docs/CampaignsApi.md#gettingallyoursenders) | **Get** /senders/find_all.{Format} | Getting All Your Senders
*CampaignsApi* | [**GettingCampaignDetails**](docs/CampaignsApi.md#gettingcampaigndetails) | **Get** /campaigns/{CampaignID}/view.{Format} | Getting Campaign Details
*CampaignsApi* | [**GettingSenderDetails**](docs/CampaignsApi.md#gettingsenderdetails) | **Get** /senders/find_one.{Format} | Getting Sender Details
*CampaignsApi* | [**LinkActivity**](docs/CampaignsApi.md#linkactivity) | **Get** /campaigns/{CampaignID}/stats/links.{Format} | Link Activity
*CampaignsApi* | [**SchedulingACampaign**](docs/CampaignsApi.md#schedulingacampaign) | **Post** /campaigns/{CampaignID}/schedule.{Format} | Scheduling A Campaign
*CampaignsApi* | [**SendingACampaign**](docs/CampaignsApi.md#sendingacampaign) | **Post** /campaigns/{CampaignID}/send.{Format} | Sending a campaign
*CampaignsApi* | [**TestingACampaign**](docs/CampaignsApi.md#testingacampaign) | **Post** /campaigns/{CampaignID}/send_test.{Format} | Testing a campaign
*CampaignsApi* | [**UnschedulingACampaign**](docs/CampaignsApi.md#unschedulingacampaign) | **Post** /campaigns/{CampaignID}/unschedule.{Format} | Unscheduling a campaign
*CampaignsApi* | [**UpdatingADraftCampaign**](docs/CampaignsApi.md#updatingadraftcampaign) | **Post** /campaigns/{CampaignID}/update.{Format} | Updating A Draft Campaign
*MailingListsApi* | [**CreatingACustomField**](docs/MailingListsApi.md#creatingacustomfield) | **Post** /lists/{MailingListID}/customfields/create.{Format} | Creating a custom field
*MailingListsApi* | [**CreatingAMailingList**](docs/MailingListsApi.md#creatingamailinglist) | **Post** /lists/create.{Format} | Creating a mailing list
*MailingListsApi* | [**DeletingAMailingList**](docs/MailingListsApi.md#deletingamailinglist) | **Delete** /lists/{MailingListID}/delete.{Format} | Deleting a mailing list
*MailingListsApi* | [**GettingAllActiveMailingLists**](docs/MailingListsApi.md#gettingallactivemailinglists) | **Get** /lists.{Format} | Getting all active mailing lists
*MailingListsApi* | [**GettingAllActiveMailingListsWithPaging**](docs/MailingListsApi.md#gettingallactivemailinglistswithpaging) | **Get** /lists/{Page}/{PageSize}.{Format} | Getting all active mailing lists with paging
*MailingListsApi* | [**GettingMailingListDetails**](docs/MailingListsApi.md#gettingmailinglistdetails) | **Get** /lists/{MailingListID}/details.{Format} | Getting mailing list details
*MailingListsApi* | [**RemovingACustomField**](docs/MailingListsApi.md#removingacustomfield) | **Delete** /lists/{MailingListID}/customfields/{CustomFieldID}/delete.{Format} | Removing a custom field
*MailingListsApi* | [**UpdatingACustomField**](docs/MailingListsApi.md#updatingacustomfield) | **Post** /lists/{MailingListID}/customfields/{CustomFieldID}/update.{Format} | Updating a custom field
*MailingListsApi* | [**UpdatingAMailingList**](docs/MailingListsApi.md#updatingamailinglist) | **Post** /lists/{MailingListID}/update.{Format} | Updating a mailing list
*SegmentsApi* | [**AddingCriteriaToSegments**](docs/SegmentsApi.md#addingcriteriatosegments) | **Post** /lists/{MailingListID}/segments/{SegmentID}/criteria/add.{Format} | Adding criteria to segments
*SegmentsApi* | [**CreatingANewSegment**](docs/SegmentsApi.md#creatinganewsegment) | **Post** /lists/{MailingListID}/segments/create.{Format} | Creating a new segment
*SegmentsApi* | [**DeletingASegment**](docs/SegmentsApi.md#deletingasegment) | **Delete** /lists/{MailingListID}/segments/{SegmentID}/delete.{Format} | Deleting A Segment
*SegmentsApi* | [**GettingSegmentDetails**](docs/SegmentsApi.md#gettingsegmentdetails) | **Get** /lists/{MailingListID}/segments/{SegmentID}/details.{Format} | Getting segment details
*SegmentsApi* | [**GettingSegmentSubscribers**](docs/SegmentsApi.md#gettingsegmentsubscribers) | **Get** /lists/{MailingListID}/segments/{SegmentID}/members.{Format} | Getting segment subscribers
*SegmentsApi* | [**GettingSegments**](docs/SegmentsApi.md#gettingsegments) | **Get** /lists/{MailingListID}/segments.{Format} | Getting segments
*SegmentsApi* | [**UpdatingASegment**](docs/SegmentsApi.md#updatingasegment) | **Post** /lists/{MailingListID}/segments/{SegmentID}/update.{Format} | Updating a segment
*SegmentsApi* | [**UpdatingSegmentCriteria**](docs/SegmentsApi.md#updatingsegmentcriteria) | **Post** /lists/{MailingListID}/segments/{SegmentID}/criteria/{CriteriaID}/update.{Format} | Updating segment criteria
*SubscribersApi* | [**AddingMultipleSubscribers**](docs/SubscribersApi.md#addingmultiplesubscribers) | **Post** /subscribers/{MailingListID}/subscribe_many.{Format} | Adding multiple subscribers
*SubscribersApi* | [**AddingSubscribers**](docs/SubscribersApi.md#addingsubscribers) | **Post** /subscribers/{MailingListID}/subscribe.{Format} | Adding subscribers
*SubscribersApi* | [**GetSubscriberByEmailAddress**](docs/SubscribersApi.md#getsubscriberbyemailaddress) | **Get** /subscribers/{MailingListID}/view.{Format} | Get subscriber by email address
*SubscribersApi* | [**GetSubscriberById**](docs/SubscribersApi.md#getsubscriberbyid) | **Get** /subscribers/{MailingListID}/find/{SubscriberID}.{Format} | Get subscriber by id
*SubscribersApi* | [**GettingSubscribers**](docs/SubscribersApi.md#gettingsubscribers) | **Get** /lists/{MailingListID}/subscribers/{Status}.{Format} | Getting subscribers
*SubscribersApi* | [**RemovingASubscriber**](docs/SubscribersApi.md#removingasubscriber) | **Post** /subscribers/{MailingListID}/remove.{Format} | Removing a subscriber
*SubscribersApi* | [**RemovingMultipleSubscribers**](docs/SubscribersApi.md#removingmultiplesubscribers) | **Post** /subscribers/{MailingListID}/remove_many.{Format} | Removing multiple subscribers
*SubscribersApi* | [**UnsubscribingSubscribersFromAccount**](docs/SubscribersApi.md#unsubscribingsubscribersfromaccount) | **Post** /subscribers/unsubscribe.{Format} | Unsubscribing subscribers from account
*SubscribersApi* | [**UnsubscribingSubscribersFromMailingList**](docs/SubscribersApi.md#unsubscribingsubscribersfrommailinglist) | **Post** /subscribers/{MailingListID}/unsubscribe.{Format} | Unsubscribing subscribers from mailing list
*SubscribersApi* | [**UnsubscribingSubscribersFromMailingListAndASpecifiedCampaign**](docs/SubscribersApi.md#unsubscribingsubscribersfrommailinglistandaspecifiedcampaign) | **Post** /subscribers/{MailingListID}/{CampaignID}/unsubscribe.{Format} | Unsubscribing subscribers from mailing list and a specified campaign
*SubscribersApi* | [**UpdatingASubscriber**](docs/SubscribersApi.md#updatingasubscriber) | **Post** /subscribers/{MailingListID}/update/{SubscriberID}.{Format} | Updating a subscriber


## Documentation For Models

 - [A](docs/A.md)
 - [AbCampaignData](docs/AbCampaignData.md)
 - [AbTestCampaignSummaryResponse](docs/AbTestCampaignSummaryResponse.md)
 - [ActivityByLocationResponse](docs/ActivityByLocationResponse.md)
 - [AddingCriteriaToSegmentsRequest](docs/AddingCriteriaToSegmentsRequest.md)
 - [AddingCriteriaToSegmentsResponse](docs/AddingCriteriaToSegmentsResponse.md)
 - [AddingMultipleSubscribersRequest](docs/AddingMultipleSubscribersRequest.md)
 - [AddingMultipleSubscribersResponse](docs/AddingMultipleSubscribersResponse.md)
 - [AddingSubscribersRequest](docs/AddingSubscribersRequest.md)
 - [AddingSubscribersResponse](docs/AddingSubscribersResponse.md)
 - [Analytic](docs/Analytic.md)
 - [B](docs/B.md)
 - [Campaign](docs/Campaign.md)
 - [CampaignSummaryResponse](docs/CampaignSummaryResponse.md)
 - [CloningAnExistingCampaignResponse](docs/CloningAnExistingCampaignResponse.md)
 - [Context](docs/Context.md)
 - [Context110](docs/Context110.md)
 - [Context118](docs/Context118.md)
 - [Context132](docs/Context132.md)
 - [Context140](docs/Context140.md)
 - [Context145](docs/Context145.md)
 - [Context148](docs/Context148.md)
 - [Context17](docs/Context17.md)
 - [Context32](docs/Context32.md)
 - [Context37](docs/Context37.md)
 - [Context52](docs/Context52.md)
 - [Context64](docs/Context64.md)
 - [Context66](docs/Context66.md)
 - [Context72](docs/Context72.md)
 - [Context84](docs/Context84.md)
 - [Context89](docs/Context89.md)
 - [Context93](docs/Context93.md)
 - [CreatingACustomFieldRequest](docs/CreatingACustomFieldRequest.md)
 - [CreatingACustomFieldResponse](docs/CreatingACustomFieldResponse.md)
 - [CreatingADraftCampaignRequest](docs/CreatingADraftCampaignRequest.md)
 - [CreatingADraftCampaignResponse](docs/CreatingADraftCampaignResponse.md)
 - [CreatingAMailingListRequest](docs/CreatingAMailingListRequest.md)
 - [CreatingAMailingListResponse](docs/CreatingAMailingListResponse.md)
 - [CreatingANewSegmentRequest](docs/CreatingANewSegmentRequest.md)
 - [CreatingANewSegmentResponse](docs/CreatingANewSegmentResponse.md)
 - [Criterion](docs/Criterion.md)
 - [CustomField](docs/CustomField.md)
 - [CustomField53](docs/CustomField53.md)
 - [CustomFieldsDefinition](docs/CustomFieldsDefinition.md)
 - [DeletingACampaignResponse](docs/DeletingACampaignResponse.md)
 - [DeletingAMailingListResponse](docs/DeletingAMailingListResponse.md)
 - [DeletingASegmentResponse](docs/DeletingASegmentResponse.md)
 - [Format](docs/Format.md)
 - [GetAllCampaignsResponse](docs/GetAllCampaignsResponse.md)
 - [GetCampaignStatisticsResponse](docs/GetCampaignStatisticsResponse.md)
 - [GetCampaignStatisticsWithPagingFilteredResponse](docs/GetCampaignStatisticsWithPagingFilteredResponse.md)
 - [GetCampaignsByPageAndPagesizeResponse](docs/GetCampaignsByPageAndPagesizeResponse.md)
 - [GetCampaignsByPageResponse](docs/GetCampaignsByPageResponse.md)
 - [GetSubscriberByEmailAddressResponse](docs/GetSubscriberByEmailAddressResponse.md)
 - [GetSubscriberByIdResponse](docs/GetSubscriberByIdResponse.md)
 - [GettingAllActiveMailingListsResponse](docs/GettingAllActiveMailingListsResponse.md)
 - [GettingAllActiveMailingListsWithPagingResponse](docs/GettingAllActiveMailingListsWithPagingResponse.md)
 - [GettingAllYourSendersResponse](docs/GettingAllYourSendersResponse.md)
 - [GettingCampaignDetailsResponse](docs/GettingCampaignDetailsResponse.md)
 - [GettingMailingListDetailsResponse](docs/GettingMailingListDetailsResponse.md)
 - [GettingSegmentDetailsResponse](docs/GettingSegmentDetailsResponse.md)
 - [GettingSegmentSubscribersResponse](docs/GettingSegmentSubscribersResponse.md)
 - [GettingSegmentsResponse](docs/GettingSegmentsResponse.md)
 - [GettingSenderDetailsResponse](docs/GettingSenderDetailsResponse.md)
 - [GettingSubscribersResponse](docs/GettingSubscribersResponse.md)
 - [ImportOperation](docs/ImportOperation.md)
 - [ImportOperation19](docs/ImportOperation19.md)
 - [LinkActivityResponse](docs/LinkActivityResponse.md)
 - [MailingList](docs/MailingList.md)
 - [MailingList68](docs/MailingList68.md)
 - [MailingList69](docs/MailingList69.md)
 - [MailingList85](docs/MailingList85.md)
 - [MailingLists](docs/MailingLists.md)
 - [MailingLists119](docs/MailingLists119.md)
 - [MailingLists134](docs/MailingLists134.md)
 - [ModelType](docs/ModelType.md)
 - [Paging](docs/Paging.md)
 - [Paging76](docs/Paging76.md)
 - [RemovingACustomFieldResponse](docs/RemovingACustomFieldResponse.md)
 - [RemovingASubscriberRequest](docs/RemovingASubscriberRequest.md)
 - [RemovingASubscriberResponse](docs/RemovingASubscriberResponse.md)
 - [RemovingMultipleSubscribersRequest](docs/RemovingMultipleSubscribersRequest.md)
 - [RemovingMultipleSubscribersResponse](docs/RemovingMultipleSubscribersResponse.md)
 - [ReplyToEmail](docs/ReplyToEmail.md)
 - [SchedulingACampaignRequest](docs/SchedulingACampaignRequest.md)
 - [SchedulingACampaignResponse](docs/SchedulingACampaignResponse.md)
 - [Segment](docs/Segment.md)
 - [Sender](docs/Sender.md)
 - [SendingACampaignResponse](docs/SendingACampaignResponse.md)
 - [ShortBy](docs/ShortBy.md)
 - [SortMethod](docs/SortMethod.md)
 - [Status](docs/Status.md)
 - [Subscriber](docs/Subscriber.md)
 - [Subscribers](docs/Subscribers.md)
 - [Subscribers150](docs/Subscribers150.md)
 - [TestingACampaignRequest](docs/TestingACampaignRequest.md)
 - [TestingACampaignResponse](docs/TestingACampaignResponse.md)
 - [UnschedulingACampaignResponse](docs/UnschedulingACampaignResponse.md)
 - [UnsubscribingSubscribersFromAccountRequest](docs/UnsubscribingSubscribersFromAccountRequest.md)
 - [UnsubscribingSubscribersFromAccountResponse](docs/UnsubscribingSubscribersFromAccountResponse.md)
 - [UnsubscribingSubscribersFromMailingListAndASpecifiedCampaignRequest](docs/UnsubscribingSubscribersFromMailingListAndASpecifiedCampaignRequest.md)
 - [UnsubscribingSubscribersFromMailingListAndASpecifiedCampaignResponse](docs/UnsubscribingSubscribersFromMailingListAndASpecifiedCampaignResponse.md)
 - [UnsubscribingSubscribersFromMailingListRequest](docs/UnsubscribingSubscribersFromMailingListRequest.md)
 - [UnsubscribingSubscribersFromMailingListResponse](docs/UnsubscribingSubscribersFromMailingListResponse.md)
 - [UpdatingACustomFieldRequest](docs/UpdatingACustomFieldRequest.md)
 - [UpdatingACustomFieldResponse](docs/UpdatingACustomFieldResponse.md)
 - [UpdatingADraftCampaignRequest](docs/UpdatingADraftCampaignRequest.md)
 - [UpdatingADraftCampaignResponse](docs/UpdatingADraftCampaignResponse.md)
 - [UpdatingAMailingListRequest](docs/UpdatingAMailingListRequest.md)
 - [UpdatingAMailingListResponse](docs/UpdatingAMailingListResponse.md)
 - [UpdatingASegmentRequest](docs/UpdatingASegmentRequest.md)
 - [UpdatingASegmentResponse](docs/UpdatingASegmentResponse.md)
 - [UpdatingASubscriberRequest](docs/UpdatingASubscriberRequest.md)
 - [UpdatingASubscriberResponse](docs/UpdatingASubscriberResponse.md)
 - [UpdatingSegmentCriteriaRequest](docs/UpdatingSegmentCriteriaRequest.md)
 - [UpdatingSegmentCriteriaResponse](docs/UpdatingSegmentCriteriaResponse.md)
 - [WithStatistics](docs/WithStatistics.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author



