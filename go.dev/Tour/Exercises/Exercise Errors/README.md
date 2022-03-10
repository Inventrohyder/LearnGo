# Exercise: Errors

Copy your `Sqrt` function from the [earlier exercise](../Exercise%20Loops%20and%20Functions/exercise-loops-and-functions.go) and modify it to return an `error` value.

`Sqrt` should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

> `type ErrNegativeSqrt float64`

and make it an `error` by giving it a

> `func (e ErrNegativeSqrt) Error() string`

method such that `ErrNegativeSqrt(-2).Error()` returns `"cannot Sqrt negative number: -2"`.

**Note:** A call to `fmt.Sprint(e)` inside the `Error` method will send the program into an infinite loop. You can avoid this by converting `e` first: `fmt.Sprint(float64(e))`. Why?

> That is because `fmt.Sprint(e)` will call `e.Error()` to convert `e` to a `string` which intern will call `fmt.Sprint(e)` again. leading to a non-ending recursive loop.

Change your `Sqrt` function to return an `ErrNegativeSqrt` value when given a negative number.
