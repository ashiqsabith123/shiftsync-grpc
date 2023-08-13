package usecases

import (
	"context"
	"errors"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
	repo "github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/repository/interfaces"
	service "github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/usecases/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	adminRepo repo.AdminRepository
}

func NewAdminUseCase(adRep repo.AdminRepository) service.AdminUseCase {
	return &AdminUseCase{adminRepo: adRep}
}

func (a *AdminUseCase) SignUp(ctx context.Context, admin domain.Admin) error {
	_, err := a.adminRepo.FindAdmin(ctx, admin)
	if err == nil {
		return errors.New("admin already exist")
	}

	hash, hasherr := bcrypt.GenerateFromPassword([]byte(admin.Password), 14)

	if hasherr != nil {
		return errors.New("hashing failed" + hasherr.Error())
	}

	admin.Password = string(hash)

	if err := a.adminRepo.SaveAdmin(ctx, admin); err != nil {
		return errors.New("unable to add admin " + err.Error())
	}

	return nil
}

func (a *AdminUseCase) SignIn(ctx context.Context, details domain.Admin) (domain.Admin, error) {
	admin, err := a.adminRepo.FindAdmin(ctx, details)
	if err != nil {
		return details, errors.New("invalid credentials " + err.Error())
	}

	if berr := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(details.Password)); berr != nil {
		return details, errors.New("incorrect password")
	}

	return admin, nil
}
