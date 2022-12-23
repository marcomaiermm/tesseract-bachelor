import qs from "qs";
import z from "zod";

/**
 * Base API function. All API Functions should be built on top of this.
 * returns infered data type from zod schema
 * @param url
 * @param options
 * @returns
 */
const baseApiCall = async <T>(
  url: string,
  options: qs.ParsedQs = {}
): Promise<T> => {
  // if url is not starting with a slash, add one
  if (!url.startsWith("/")) url = "/" + url;
  const response = await fetch(
    `${import.meta.env.VITE_API_URL}${url}?${qs.stringify(options)}`
  );
  const data = await response.json();
  return data;
};

export default baseApiCall;
