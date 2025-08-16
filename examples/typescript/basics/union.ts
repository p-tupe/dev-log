// This code snippet shows how to create a union type. Useful for mapping API responses.
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
