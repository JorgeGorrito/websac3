package seeders

import (
	"strconv"
	"websac3/adapter/out/persistence/postgresql/db/postgres"
	"websac3/adapter/out/persistence/postgresql/dto"
	models2 "websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"

	"gorm.io/gorm"
)

const (
	DEFAULT_PATH_HIGHER_EDUCATION_INSTITUTION_SEED = "adapter/out/persistence/seeders/seeds/higher_education_institutions.json"
)

type higherEducationInstitutions struct{}

func HigherEducationInstitutions() Seeder {
	return &higherEducationInstitutions{}
}

func (h *higherEducationInstitutions) getSniesAsUint(snies string) (uint, error) {
	var sniesUint uint64
	sniesUint, err := strconv.ParseUint(snies, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(sniesUint), nil
}

func (h *higherEducationInstitutions) getOwnershipId(transaction *gorm.DB, ownershipName string) (uint, error) {
	var ownership uint
	result := transaction.Model(&models2.Ownership{}).Select("id").Where("name = ?", ownershipName).First(&ownership)
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return ownership, result.Error
}

func (h *higherEducationInstitutions) getInstitutionalCategoryId(transaction *gorm.DB, institutionalCategoryName string) (uint, error) {
	var institutionalCategory uint
	result := transaction.Model(&models2.InstitutionalCategory{}).Select("id").Where("name = ?", institutionalCategoryName).First(&institutionalCategory)
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return institutionalCategory, result.Error
}

func (h *higherEducationInstitutions) getMunicipalityId(transaction *gorm.DB, municipalityName string) (uint, error) {
	var municipality uint
	result := transaction.Model(&models2.Municipality{}).Select("id").Where("name = ?", municipalityName).First(&municipality)
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return municipality, result.Error
}

func (h *higherEducationInstitutions) Seed(tx persistence.Transaction) error {
	var pgTx *postgres.Transaction = tx.(*postgres.Transaction)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []dto.HigherEducationInstitution = make([]dto.HigherEducationInstitution, 0)
	if err := decoder.Decode(DEFAULT_PATH_HIGHER_EDUCATION_INSTITUTION_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, higherEducationInstitution := range dataToSeed {
		snies, err := h.getSniesAsUint(higherEducationInstitution.Snies)
		if err != nil {
			return err
		}

		ownershipId, err := h.getOwnershipId(pgTx.Tx(), higherEducationInstitution.Ownership)
		if err != nil {
			return err
		}

		institutionalCategoryId, err := h.getInstitutionalCategoryId(pgTx.Tx(), higherEducationInstitution.InstitutionalCategory)
		if err != nil {
			return err
		}

		municipalityId, err := h.getMunicipalityId(pgTx.Tx(), higherEducationInstitution.Municipality)
		if err != nil {
			return err
		}

		departmentId, err := func() (uint, error) {
			var department uint
			result := pgTx.Tx().Model(&models2.Department{}).Select("id").Where("name = ?", higherEducationInstitution.Department).First(&department)
			if result.RowsAffected == 0 {
				return 0, nil
			}
			return department, result.Error
		}()
		if err != nil {
			return err
		}

		if err := pgTx.Tx().Create(&models2.HigherEducationInstitution{
			Snies:                   snies,
			SniesParent:             nil,
			Name:                    higherEducationInstitution.Name,
			OwnershipID:             ownershipId,
			InstitutionalCategoryID: institutionalCategoryId,
			MunicipalityID:          municipalityId,
			DepartmentID:            departmentId,
		}).Error; err != nil {
			return err
		}
	}

	for _, higherEducationInstitution := range dataToSeed {
		sniesParent, err := h.getSniesAsUint(higherEducationInstitution.SniesParent)
		if err != nil {
			return err
		}
		snies, err := h.getSniesAsUint(higherEducationInstitution.Snies)
		if err != nil {
			return err
		}
		if err := pgTx.Tx().Model(&models2.HigherEducationInstitution{}).
			Where("snies = ?", snies).
			Update("snies_parent", sniesParent).Error; err != nil {
			return err
		}
	}
	return nil
}
