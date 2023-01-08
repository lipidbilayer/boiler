import { Button, Nav, NavItem } from "reactstrap";
import Logo from "../../logo/Logo";
import Link from "next/link";
import { useRouter } from "next/router";

const navigation = [
  {
    title: "Reporting",
    href: "/",
    icon: "bi bi-speedometer2",
  },
  {
    title: "Scan Armada",
    href: "/scanner",
    icon: "bi bi-qr-code-scan",
  },
  {
    title: "Data Pengguna",
    href: "/management/users",
    icon: "bi bi-people",
  },
  {
    title: "Data Lokasi",
    href: "/management/locations",
    icon: "bi bi-geo",
  },
  {
    title: "Data Armada",
    href: "/management/vehicles",
    icon: "bi bi-truck-flatbed",
  },
  {
    title: "Badges",
    href: "/ui/badges",
    icon: "bi bi-patch-check",
  },
  {
    title: "Buttons",
    href: "/ui/buttons",
    icon: "bi bi-hdd-stack",
  },
  {
    title: "Cards",
    href: "/ui/cards",
    icon: "bi bi-card-text",
  },
  {
    title: "Grid",
    href: "/ui/grid",
    icon: "bi bi-columns",
  },
  {
    title: "Table",
    href: "/ui/tables",
    icon: "bi bi-layout-split",
  },
  {
    title: "Forms",
    href: "/ui/forms",
    icon: "bi bi-textarea-resize",
  },
  {
    title: "Breadcrumbs",
    href: "/ui/breadcrumbs",
    icon: "bi bi-link",
  },
  {
    title: "About",
    href: "/about",
    icon: "bi bi-people",
  },
];

const Sidebar = ({ showMobilemenu, loading }) => {
  let curl = useRouter();
  const location = curl.pathname;

  const SidebarMenu = () => {
      return (        
        <Nav vertical className="sidebarNav">
        {navigation.map((navi, index) => (
            <NavItem key={index} className="sidenav-bg">
              <Link href={navi.href}>
                <a
                  className={
                    location === navi.href
                      ? "text-primary nav-link py-3"
                      : "nav-link text-secondary py-3"
                  }
                >
                  <i className={navi.icon}></i>
                  <span className="ms-3 d-inline-block">{navi.title}</span>
                </a>
              </Link>
            </NavItem>
          ))}
        </Nav>)
  }

  return (
    <div className="p-3">
      <div className="d-flex align-items-center">
        <Logo />
        <Button
          close
          size="sm"
          className="ms-auto d-lg-none"
          onClick={showMobilemenu}
        ></Button>
      </div>
      <div className="pt-4 mt-2">
        <SidebarMenu></SidebarMenu>
      </div>
    </div>
  );
};

export default Sidebar;