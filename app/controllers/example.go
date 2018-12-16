package controllers

import (
	"net/http"

	"github.com/binatify/gin-template/app/models"
	"github.com/binatify/gin-template/base/context"
)

var (
	Example *_Example
)

type _Example struct{}

func (_ *_Example) Create(ctx *context.Context) {
	var input *CreateExampleInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Logger().Errorf("Unmarshal json error: %v", err)
		ctx.JSON(http.StatusBadRequest, "Please input right parameters")
		return
	}

	example := models.NewExampleModel(input.Name)
	example.Phone = input.Phone
	if err := example.Save(); err != nil {
		ctx.Logger().Errorf("exmaple.Save(): %v", err)
		ctx.JSON(http.StatusBadRequest, "Invalid Parameter")
		return
	}

	ctx.Logger().Infof("exmaple.Save(%v): success.", example.ID.Hex())
	ctx.JSON(http.StatusOK, NewShowExampleOutput(example))
}

func (_ *_Example) Update(ctx *context.Context) {
	id := ctx.Param("id")
	example, err := models.Example.Find(id)
	if err != nil {
		ctx.Logger().Errorf("models.Example.Find(%v): %v", id, err)
		ctx.JSON(http.StatusBadRequest, "Please input right id")
		return
	}

	var input *UpdateExampleInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Logger().Errorf("Unmarshal json error: %v", err)
		ctx.JSON(http.StatusBadRequest, "Please input right parameters")
		return
	}

	example.Name = input.Name
	if err := example.Save(); err != nil {
		ctx.Logger().Errorf("example.Save(): %v", err)
		ctx.JSON(http.StatusInternalServerError, "Internal error")
		return
	}

	ctx.JSON(http.StatusOK, input)
}

func (_ *_Example) Show(ctx *context.Context) {
	id := ctx.Param("id")
	example, err := models.Example.Find(id)
	if err != nil {
		ctx.Logger().Errorf("models.Example.Find(%v): %v", id, err)
		ctx.JSON(http.StatusBadRequest, "Please input right id")
		return
	}

	ctx.JSON(http.StatusOK, NewShowExampleOutput(example))
}

func (_ *_Example) All(ctx *context.Context) {
	var input ListExamplesInput // DO NOT use pointer
	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.Logger().Errorf("ctx.BindQuery(%v): %v", input, err)
		ctx.JSON(http.StatusBadRequest, "Wrong params.")
		return
	}

	examples, err := models.Example.List(100)
	if err != nil {
		ctx.Logger().Errorf("models.Example.List(): %v", err)
		ctx.JSON(http.StatusBadRequest, "Something wrong happened when listing the examples.")
		return
	}

	res := make([]*ShowExampleOutput, 0)
	for _, v := range examples {
		res = append(res, NewShowExampleOutput(v))
	}

	ctx.JSON(http.StatusOK, res)
}

func (_ *_Example) Remove(ctx *context.Context) {
	id := ctx.Param("id")

	example, err := models.Example.Find(id)
	if err != nil {
		ctx.Logger().Errorf("models.Example.Find(%v): %v", id, err)
		ctx.JSON(http.StatusBadRequest, "Please input right id")
		return
	}

	if err := example.Delete(); err != nil {
		ctx.Logger().Errorf("example.Delete(): %v", err)
		ctx.JSON(http.StatusBadRequest, "Delete example failed.")
		return
	}

	ctx.JSON(http.StatusOK, "Successfully delete object.")
}
