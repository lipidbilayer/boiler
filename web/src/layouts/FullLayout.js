import React, { useEffect, useRef } from "react";
import { Container } from "reactstrap";
// import UserProfile from "../api/user/user";
import Header from "./header/Header";
import Sidebar from "./sidebars/vertical/Sidebar";

const FullLayout = ({ children }) => {
  const [open, setOpen] = React.useState(false);
  const [loading, setLoading] = React.useState(false);
  const [user, setUser] = React.useState(null)
  const sidebarRef = React.useRef(null)
  const showMobilemenu = () => {
    setOpen(!open);
  };

  const test = () => {
    if(open) setOpen(false)
  }

  useEffect(() => {
    // UserProfile({setLoading: setLoading, setData: setUser})
    document.addEventListener("mousedown", handleClickOutside, false);
    return () => document.removeEventListener('mousedown', handleClickOutside, false);
  }, [])

  const handleClickOutside = event => {
    if (sidebarRef.current && !sidebarRef.current.contains(event.target)) {
      setOpen(false)
    }
  };

  return (
    <main>
      <div className="pageWrapper d-md-block d-lg-flex">

        <aside ref={sidebarRef}
          className={`sidebarArea shadow bg-white ${
            !open ? "" : "showSidebar"
          }`}
        >
          <Sidebar loading={loading} user={user} showMobilemenu={() => showMobilemenu()}/>
        </aside>
        <div className="contentArea">
          <Header showMobmenu={() => showMobilemenu()} />

          <Container className="p-4 wrapper" fluid>
            <div>{children}</div>
          </Container>
        </div>
      </div>
    </main>
  );
};

export default FullLayout;