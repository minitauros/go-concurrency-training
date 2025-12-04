package solution

import (
	"context"
	"errors"
	"strings"
	"testing"

	mocks2 "github.com/minitauros/go-concurrency-training/courses/testing/5_mocks"
	"github.com/minitauros/go-concurrency-training/courses/testing/5_mocks/mocks"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_CreateUserUseCase_CreateUser(t *testing.T) {
	Convey("CreateUserUseCase.CreateUser()", t, func() {
		ctx := context.Background()
		userRepo := mocks.NewMockUserRepository(t)
		uc := &mocks2.CreateUserUseCase{
			Repo: userRepo,
		}

		Convey("If the username is not allowed", func() {
			username := "disallowed"

			Convey("Returns error", func() {
				id, err := uc.CreateUser(ctx, username)

				So(err, ShouldEqual, mocks2.ErrUsernameNotAllowed)
				So(id, ShouldEqual, 0)
			})
		})
		Convey("If the username is too long", func() {
			username := strings.Repeat("a", 21)

			Convey("Returns error", func() {
				id, err := uc.CreateUser(ctx, username)

				So(err, ShouldEqual, mocks2.ErrUsernameTooLong)
				So(id, ShouldEqual, 0)
			})
		})

		Convey("If the username is allowed", func() {
			Convey("If the username is at the maximum allowed length", func() {
				username := strings.Repeat("a", 20)

				Convey("If the user cannot be created in the db", func() {
					expectedErr := errors.New("something went wrong")
					user := mocks2.User{
						Username: username,
					}
					userRepo.EXPECT().CreateUser(ctx, user).Return(0, expectedErr).Once()

					Convey("Returns error", func() {
						id, err := uc.CreateUser(ctx, username)

						So(err, ShouldEqual, expectedErr)
						So(id, ShouldEqual, 0)
					})
				})

				Convey("If the user can be created in the db", func() {
					expectedUserID := int64(123)
					userRepo.EXPECT().
						CreateUser(ctx, mocks2.User{Username: username}).
						Return(expectedUserID, nil).Once()

					Convey("Returns no error", func() {
						id, err := uc.CreateUser(ctx, username)

						So(err, ShouldBeNil)
						So(id, ShouldEqual, expectedUserID)
					})
				})
			})
		})
	})
}
