package repositories

import (
	"errors"
	"testing"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoSuite struct {
	suite.Suite
	mockCollection *mocks.MongoCollection
	mockUser       domain.User
	mockEmptyUser  domain.User
	mockUserID     primitive.ObjectID
}

func (suite *RepoSuite) SetupTest() {
	suite.mockCollection = mocks.NewMongoCollection(suite.T())

	suite.mockUser = domain.User{
		ID:       primitive.NewObjectID(),
		Username: "username",
		Role:     "admin",
		Password: "password",
	}

	suite.mockEmptyUser = domain.User{}
	suite.mockUserID = primitive.NewObjectID()
}

func (suite *RepoSuite) TestSuccessCreate() {

	suite.mockCollection.On("InsertOne", mock.Anything, mock.AnythingOfType("domain.User")).Return(&mongo.InsertOneResult{
		InsertedID: suite.mockUserID,
	}, nil).Once()

	ur := NewUserRepository(suite.mockCollection)

	usr, err := ur.Save(suite.mockUser)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), suite.mockUserID, usr.ID)

	suite.mockCollection.AssertExpectations(suite.T())
}

func (suite *RepoSuite) TestErrorCreate() {
	suite.mockCollection.On("InsertOne", mock.Anything, mock.AnythingOfType("domain.User")).Return(&mongo.InsertOneResult{
		InsertedID: suite.mockUserID,
	}, errors.New("error")).Once()

	ur := NewUserRepository(suite.mockCollection)

	usr, err := ur.Save(suite.mockEmptyUser)

	assert.Error(suite.T(), err)

	assert.Equal(suite.T(), suite.mockEmptyUser, usr)

	suite.mockCollection.AssertExpectations(suite.T())
}

// func (suite *RepoSuite) TestSuccessGetByID() {
// 	suite.mockCollection.On("FindOne", mock.Anything, mock.AnythingOfType("primitive.M")).Return(&mongo.SingleResult{}, nil).Once()

// 	ur := NewUserRepository(suite.mockCollection)

// 	usr, err := ur.FindUser(suite.mockUserID.Hex())

// 	assert.NoError(suite.T(), err)

// 	assert.IsType(suite.T(), suite.mockUser, usr)

// 	suite.mockCollection.AssertExpectations(suite.T())
// }

func (suite *RepoSuite) TestErrorGetByID() {
	suite.mockCollection.On("FindOne", mock.Anything, mock.AnythingOfType("primitive.M")).Return(&mongo.SingleResult{}, nil).Once()

	ur := NewUserRepository(suite.mockCollection)

	usr, err := ur.FindUser(suite.mockUserID.Hex())

	assert.Error(suite.T(), err)

	assert.Equal(suite.T(), suite.mockEmptyUser, usr)

	suite.mockCollection.AssertExpectations(suite.T())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RepoSuite))
}
