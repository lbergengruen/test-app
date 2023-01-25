package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go"
	"strconv"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	uid := strconv.FormatInt(time.Now().UnixMilli(), 10)

	compose, err := tc.NewDockerCompose("./testresources/docker-compose.yml")
	assert.NoError(t, err, "NewDockerComposeAPI()")

	t.Cleanup(func() {
		assert.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	err = compose.WithEnv(map[string]string{"uid": uid}).Up(ctx, tc.Wait(true))
	assert.NoError(t, err, "compose.Up()")
}
