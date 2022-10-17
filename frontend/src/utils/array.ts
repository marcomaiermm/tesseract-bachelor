/**
 * Returns a Map of the given array, grouped by one or more keys.
 * @param array The array to group.
 * @param keys The keys to group by.
 * @returns A Map of the given array, grouped by one or more keys.
 */
export const groupBy = <T, K extends keyof T>(
  array: T[],
  ...keys: K[]
): Map<T[K], T[]> => {
  return array.reduce((map, item) => {
    const key = keys.map((k) => item[k]).join('::') as T[K];
    const collection = map.get(key);
    if (collection) {
      collection.push(item);
    } else {
      map.set(key, [item]);
    }
    return map;
  }, new Map<T[K], T[]>());
};
