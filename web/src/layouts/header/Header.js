import React from "react";
import Link from "next/link";
import Image from "next/image";
import {
  Navbar,
  Collapse,
  Nav,
  NavItem,
  NavbarBrand,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
  Dropdown,
  Button,
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter,
} from "reactstrap";
import LogoWhite from "../../assets/images/logos/xtremelogowhite.svg";
import user1 from "../../assets/images/users/user1.jpg";

const Header = ({ showMobmenu }) => {
  const [isOpen, setIsOpen] = React.useState(false);
  const [dropdownOpen, setDropdownOpen] = React.useState(false);
  const [isLogoutOpen, setIsLogoutOpen] = React.useState(false)

  const toggleLogoutModal = () => {
    setIsLogoutOpen(!isLogoutOpen)
  }

  const toggle = () => setDropdownOpen((prevState) => !prevState);
  const Handletoggle = () => {
    setIsOpen(!isOpen);
  };

  return (
    <Navbar color="primary" dark expand="md">
      <div className="d-flex align-items-center">
        <Button color="primary" className="d-lg-none" onClick={showMobmenu}>
          <i className="bi bi-list"></i>
        </Button>
        <NavbarBrand href="/" className="d-lg-none">
          {/* <Image src={LogoWhite} alt="logo" /> */}
        </NavbarBrand>
      </div>

        <Nav className="me-auto" navbar>
        </Nav>
        {/* <Dropdown isOpen={dropdownOpen} toggle={toggle}>
          <DropdownToggle color="primary" style={{marginRight: "10px"}}>
            <div style={{ lineHeight: "0px" }}>
              <span className="bi bi-bell rounded-circle"></span>
              <Image
                src={user1}
                alt="profile"
                className="rounded-circle"
                width="30"
                height="30"
              />
            </div>
          </DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>Menu</DropdownItem>
            <DropdownItem><Link href="/notifikasi/"><a className="nav-link text-secondary">Notifikasi</a></Link></DropdownItem>
          </DropdownMenu>
        </Dropdown> */}
        
        <Link  href="/notifikasi/">
        <a>
        <Button color="primary" size="sm" className="" style={{marginRight: "10px"}}>  <i className="bi bi-bell"></i></Button></a>
        </Link>
        <Button color="warning" size="sm" className="" onClick={toggleLogoutModal}>  <i className="bi bi-box-arrow-right"></i>   Logout </Button>

        <Modal isOpen={isLogoutOpen} size="sm">
          <ModalHeader>Konfirmasi keluar dari user?</ModalHeader>
          <ModalFooter>
          <Button color="secondary" onClick={toggleLogoutModal}>
              Batal
            </Button>
            <Link href="/logout">
              <a>
              <Button color="warning">
                Submit
              </Button>
              </a>
            </Link>
          </ModalFooter>
        </Modal>
    </Navbar>
  );
};

export default Header;