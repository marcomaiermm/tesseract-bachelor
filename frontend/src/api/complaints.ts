import baseApiCall from './base';
import z from 'zod';
import { ParsedQs } from 'qs';

export const Complaint = z.object({
  id: z.number(),
  date: z.string(),
  order: z.string(),
  machine: z.string(),
  quantity: z.number().int(),
  cost: z.number(),
  intern: z.boolean(),
  reason: z.string(),
  material: z.string(),
});

type Complaint = z.infer<typeof Complaint>;

/**
 * getComlpaints - get all comlplaints from the backend api route
 * @param options
 * @returns Complaints array data
 */
export const getComplaints = async (options: ParsedQs = {}) => {
  const data = await baseApiCall<Complaint[]>('/complaint', options);
  return Complaint.array().parse(data);
};

export const getComplaint = async (id: number) => {
  const data = await baseApiCall<Complaint>(`/complaint/${id}`);
  return Complaint.parse(data);
};
