// This code snippet shows how to create a union type.
//
// X will either be {a: true, b: number} or {a: false | undefined | null, c: number}

type X =
  | {
      a: true;
      b: number;
    }
  | {
      a?: false;
      c: number;
    };

const x1: X = { a: true, b: 1 };
const x2: X = { c: 1 };

// Another example, an API where we get either an Error or Sucess.
// If it's a Sucess, then we get "data", else we get an "error" message.

type SuccessResponse = { status: "OK"; data: Record<string, any> };

type ErrorResponse = { status: string; error: string };

type APIResponse = SuccessResponse | ErrorResponse;
