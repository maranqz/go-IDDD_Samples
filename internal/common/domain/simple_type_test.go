package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type someIntType int
type someIntType64 int64

func Test_commonInt(t *testing.T) {
	zero := int64(0)

	t.Run("zero", func(t *testing.T) {
		zeroInt64 := someIntType64(0)
		got, err := newIntID(zeroInt64)

		require.Equal(t, got, zero)
		require.ErrorIs(t, err, errIntIDEmpty)
	})

	t.Run("negative", func(t *testing.T) {
		negativeInt := someIntType(-1)
		got, err := newIntID(negativeInt)

		require.Equal(t, got, zero)
		require.ErrorIs(t, err, errIntIDInvalid)
	})

	t.Run("positive", func(t *testing.T) {
		positiveInt := 11
		got, err := newIntID(positiveInt)

		require.Equal(t, got, int64(11))
		require.NoError(t, err)
	})
}

func Test_NewUserIDPtr(t *testing.T) {
	zero := UserID(0)

	t.Run("nil", func(t *testing.T) {
		var nilInt64 *UserID

		got, err := NewUserIDPtr(nilInt64)

		require.Equal(t, got, zero)
		require.ErrorIs(t, err, ErrUserIDEmpty)
	})

	t.Run("zero", func(t *testing.T) {
		zeroInt64 := someIntType64(0)
		got, err := NewUserIDPtr(&zeroInt64)

		require.Equal(t, got, zero)
		require.ErrorIs(t, err, ErrUserIDEmpty)
	})

	t.Run("negative", func(t *testing.T) {
		negativeInt := someIntType(-1)
		got, err := NewUserIDPtr(&negativeInt)

		require.Equal(t, got, zero)
		require.ErrorIs(t, err, ErrUserIDInvalid)
	})

	t.Run("positive", func(t *testing.T) {
		positiveInt := 11
		got, err := NewUserIDPtr(&positiveInt)

		require.Equal(t, got, UserID(11))
		require.NoError(t, err)
	})
}

func Test_NewUserIDFromCommonInt(t *testing.T) {
	zero := UserID(0)

	t.Run("zero", func(t *testing.T) {
		zeroInt64 := UserID(0)
		got, err := NewUserIDFromCommonInt(zeroInt64)

		require.Equal(t, got, zero)
		require.ErrorIs(t, err, ErrUserIDEmpty)
	})

	t.Run("negative", func(t *testing.T) {
		negativeInt := someIntType(-1)
		got, err := NewUserIDFromCommonInt(negativeInt)

		require.Equal(t, got, zero)
		require.ErrorIs(t, err, ErrUserIDInvalid)
	})

	t.Run("positive", func(t *testing.T) {
		positiveInt := 11
		got, err := NewUserIDFromCommonInt(positiveInt)

		require.Equal(t, got, UserID(11))
		require.NoError(t, err)
	})
}
