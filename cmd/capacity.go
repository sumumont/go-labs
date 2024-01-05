package main

import (
	"errors"
	"github.com/samber/lo"
	"os"
	"path/filepath"
)

var (
	MODEL_ABILITY_INFER = "infer"
	MODEL_ABILITY_TRAIN = "train"
)

type Capacity []string

var allCapacity = []string{"train", "infer"}

func (rec Capacity) Infer() bool {
	for _, item := range rec {
		if item == MODEL_ABILITY_INFER {
			return true
		}
	}
	return false
}
func (rec Capacity) None() bool {
	if len(rec) == 0 {
		return true
	}
	return false
}
func (rec Capacity) Train() bool {
	for _, item := range rec {
		if item == MODEL_ABILITY_TRAIN {
			return true
		}
	}
	return false
}

func (rec Capacity) OnlyInfer() bool {
	for _, item := range rec {
		if item != MODEL_ABILITY_INFER {
			return false
		}
	}
	return true
}

func (rec Capacity) OnlyTrain() bool {
	for _, item := range rec {
		if item != MODEL_ABILITY_TRAIN {
			return false
		}
	}
	return true
}

func (rec Capacity) AllCap() bool {
	o := lo.Uniq[string](rec)
	if len(o) != len(allCapacity) {
		return false
	}
	for _, item := range o {
		if !lo.Contains[string](allCapacity, item) {
			return false
		}
	}
	return true
}

func (rec Capacity) ValidateTrain(localPath string) error {
	if !rec.Train() {
		return nil
	}
	codePath := filepath.Join(localPath, "code")
	_, codeErr := os.Stat(codePath)
	if codeErr != nil {
		return errors.New("code directory not found")
	}
	return nil
}

func (rec Capacity) ValidateInfer(localPath string) error {
	if !rec.Infer() {
		return nil
	}
	codePath := filepath.Join(localPath, "infer")
	_, codeErr := os.Stat(codePath)
	if codeErr != nil {
		return errors.New("infer directory not found")
	}
	return nil
}

func (rec Capacity) Validate(localPath string) error {
	if err := rec.ValidateTrain(localPath); err != nil {
		return err
	}
	if err := rec.ValidateInfer(localPath); err != nil {
		return err
	}
	return nil
}
