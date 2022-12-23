import { ParsedQs } from "qs";
import { useQuery } from "@tanstack/react-query";
import { getComplaint, getComplaints } from "@api/complaints";

export const useComplaints = (options: ParsedQs = {}, enabled = false) => {
  return useQuery(
    ["complaints", options],
    async () => await getComplaints(options),
    { enabled }
  );
};

export const useComplaint = (id: number) => {
  return useQuery(["complaint", id], async () => await getComplaint(id));
};
