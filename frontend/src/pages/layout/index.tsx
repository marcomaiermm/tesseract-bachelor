import { Outlet } from "@tanstack/react-router";
import Footer from "./components/Footer";
import SideMenu from "./components/SideMenu";

const Layout = () => {
  return (
    <div className="w-full bg-slate-800 text-slate-100">
      <div className="flex w-full">
        <SideMenu />
        <main className="min-h-screen p-4">
          <Outlet />
        </main>
      </div>
      <Footer />
    </div>
  );
};

export default Layout;
