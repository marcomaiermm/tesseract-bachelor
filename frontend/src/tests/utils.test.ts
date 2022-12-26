import { describe, test, expect } from "vitest";
import { faker } from "@faker-js/faker";
import { groupBy } from "@utils/array";
import type { Complaint } from "@api/complaints";

const reasonse = Array.from({ length: 10 }, () =>
  faker.random.alphaNumeric(10)
);
const material = Array.from({ length: 10 }, () =>
  faker.datatype.number().toString()
);
const machine = Array.from({ length: 10 }, () =>
  faker.datatype.number().toString()
);
const order = Array.from({ length: 10 }, () =>
  faker.datatype.number().toString()
);

const complaints: Complaint[] = Array.from({ length: 100 }, () => ({
  id: faker.helpers.unique(faker.datatype.number),
  date: faker.date.past().toISOString(),
  order: order[faker.datatype.number({ min: 0, max: order.length - 1 })],
  machine: machine[faker.datatype.number({ min: 0, max: machine.length - 1 })],
  quantity: faker.datatype.number(),
  cost: faker.datatype.number(),
  intern: faker.datatype.boolean(),
  // choose random
  reason: reasonse[faker.datatype.number({ min: 0, max: reasonse.length - 1 })],
  material:
    material[faker.datatype.number({ min: 0, max: material.length - 1 })],
}));

describe("groupBy", () => {
  test("should group by reason", () => {
    const grouped = groupBy(complaints, "reason");
    // grouped should be a map with the reason as key
    expect(grouped).toBeInstanceOf(Map);
    // grouped should have the same number of keys as the number of unique reasons
    expect(grouped.size).toBe(reasonse.length);
    // grouped should have the same number of values as the number of complaints
    expect(Array.from(grouped.values()).flat().length).toBe(complaints.length);
  });
});
