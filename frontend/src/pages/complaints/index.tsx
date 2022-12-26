import { useComplaints } from "@hooks/api";
import Table from "./components/table";

const Complaints: React.FC = () => {
  const { data: complaints } = useComplaints({});

  return (
    <>
      <div>
        <div>
          <Table tableData={complaints ?? []} />
        </div>
      </div>
    </>
  );
};

export default Complaints;
