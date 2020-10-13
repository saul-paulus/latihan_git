const kodepos = require("../src/index");

describe("Search non-sense keyword", () => {
  test("Should return default results object", () => {
    const input = "asdfghjkl";
    const output = {
      count: 0,
      results: [],
    };

    expect(kodepos.search(input)).toEqual(output);
  });
});

describe("Search valid keyword", () => {
  test("Should return valid results object", () => {
    const input = "bandung";
    const output = {
      count: 0,
      results: [],
    };

    expect(kodepos.search(input)).not.toEqual(output);
  });
});

describe("Search valid postal code", () => {
  test("Should return valid results object", () => {
    const input = "40522";
    const output = {
      count: 0,
      results: [],
    };

    expect(kodepos.search(input)).not.toEqual(output);
  });
});

describe("Search with query using other than string", () => {
  test("Should throw TypeError", () => {
    try {
      kodepos.search(123);
    } catch (error) {
      expect(error).toEqual(new TypeError("Param query must be a string"));
    }
  });
});

describe("Search with limit using other than number", () => {
  test("Should throw TypeError", () => {
    try {
      kodepos.search("Bandung", "1000");
    } catch (error) {
      expect(error).toEqual(new TypeError("Param limit must be a number"));
    }
  });
});
