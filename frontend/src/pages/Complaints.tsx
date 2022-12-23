import { FC, useEffect, useState } from "react";
import { useComplaints } from "@hooks/api";

import Table from "@/components/Table";

const Complaints: FC = () => {
  const [isComplaintsFetching, setIsComplaintsFetching] = useState(false);
  const { data: complaints, status } = useComplaints({}, isComplaintsFetching);

  useEffect(() => {
    if (status === "success") {
      setIsComplaintsFetching(false);
    }
  }, [status]);

  return (
    <>
      <div>
        <button
          onClick={() => setIsComplaintsFetching(true)}
          className="transform rounded-md bg-blue-600 px-6 py-2 font-medium capitalize tracking-wide text-white transition-colors duration-300 hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
        >
          FETCH
        </button>
        <div>
          <Table
            rowsPerPage={15}
            data={complaints ?? []}
            columnHeaders={[
              "date",
              "order",
              "machine",
              "quantity",
              "cost",
              "intern",
              "reason",
              "material",
            ]}
          />
        </div>
      </div>
    </>
  );
};

export default Complaints;
