import { FC } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

import { getComplaints } from '../api/complaints';

const Complaints: FC = () => {
  const queryClient = useQueryClient();
  const query = useQuery(['complaints'], async () => getComplaints);
  return (
    <>
      <div>
        <button className="px-6 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-md hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80">
          FETCH
        </button>
      </div>
    </>
  );
};

export default Complaints;
