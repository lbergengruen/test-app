package main

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	tccompose "github.com/testcontainers/testcontainers-go/modules/compose"
)

func TestMain(t *testing.T) {
	uid := strconv.FormatInt(time.Now().UnixMilli(), 10)

	compose, err := tccompose.NewDockerCompose("./testresources/docker-compose.yml")
	assert.NoError(t, err, "NewDockerComposeAPI()")

	t.Cleanup(func() {
		assert.NoError(t, compose.Down(context.Background(), tccompose.RemoveOrphans(true), tccompose.RemoveImagesLocal), "compose.Down()")
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	err = compose.WithEnv(map[string]string{"uid": uid}).Up(ctx, tccompose.Wait(true))
	assert.NoError(t, err, "compose.Up()")
}
