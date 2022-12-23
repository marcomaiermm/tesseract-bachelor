import { FC } from "react";

type Props = {};

const Footer: FC<Props> = () => {
  return (
    <>
      <footer className="w-full bg-slate-700">
        <div className="flex flex-col items-center justify-center py-4">
          <p className="text-sm text-white">© 2022 TU Wien | Marco Maier</p>
        </div>
      </footer>
    </>
  );
};

export default Footer;
