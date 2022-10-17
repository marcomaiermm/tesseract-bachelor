import baseApiCall from './base';
import z from 'zod';

export const Defect = z.object({
  week: z.number().int().positive(),
  year: z.number().int().positive(),
  ppm: z.number().positive(),
  ppmPercent: z.number().positive(),
  total: z.number().int().positive(),
  feature: z.string(),
});

export const DefectItem = Defect.omit({ year: true }).extend({
  week: z.string(),
});

type Defect = z.infer<typeof Defect>;
type DefectItem = z.infer<typeof DefectItem>;

/**
 * getComlpaints - get all comlplaints from the backend api route
 * @param endDate
 * @returns Complaints array data
 */
export const getDefects = async (endDate: Date = new Date()) => {
  // parse endDate into format YYYY-MM-DD
  const date = endDate ? endDate.toISOString().split('T')[0] : '';
  try {
    const response = await baseApiCall<Defect[]>('/defect', { date });
    // asynchoronously parse the response as an array of Defects
    const data = z.array(Defect).parse(response);

    // generate a new array of objects with DefectItem shape. Weeks should now be a string of format YYYY"W"WW
    const defectItems = z.array(DefectItem).parse(
      data.map((item) => {
        const newItem = {
          ...item,
          week: `${item.year}W${item.week}`,
        };
        return newItem;
      })
    );
    return defectItems;
  } catch (error) {
    console.log(error);
    return [];
  }
};
