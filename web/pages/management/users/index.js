import { useRouter } from 'next/router'
import Link from 'next/link';
import React, { useState, useEffect } from 'react';
import { Row, Col, Modal, ModalHeader, ModalBody, ModalFooter, Button } from "reactstrap";
import UserTable from "../../../src/components/data/GenericTable";
import ButtonSubmitLoader from '../../../src/components/data/ButtonSubmitLoader';
import Delete from '../../../src/components/modal/delete';
import { getCookie } from 'cookies-next';
import UserList from '../../../src/api/user/admin';


export const getServerSideProps = async (ctx) => {
  const header = {
    title: "Management User",
    subtitle: "",
    head: ["Nama", "Username", "Action"],
    show: ["name", "username", "action"],
    rows: []
    };
  const actions = [
      {
        icon: "bi bi-pencil",
        color: "warning",
        url: "/management/users/",
        text: "Edit"
      },
      {
        icon: "bi bi-trash",
        color: "danger",
        text: "Delete"
      }
    ];
  return {
    props: {header: header, actions: actions}
  }
}

const Tables = ({header, actions}) => {

  const router = useRouter()
  const [data, setData] = useState([])
  const [deleteModal, setDeleteModal] = useState(false)
  const [loader, setLoader] = useState(false)
  const [deleteObject, setDeleteObject] = useState({id: 0, name:""})

  useEffect(() => {
    UserList({setData: setData, actions: actions, setLoading: null})
  }, [UserList])

  const deleteToggle = ({user}) => {
    const objUser = user ? user : {id: 0, name:""}
    setDeleteObject(objUser);
    setDeleteModal(!deleteModal);
    
  }

  const deleteUser = async (event) => {
    setLoader(true)
    event.preventDefault()
    const endpoint = process.env.NEXT_PUBLIC_SERVER_API_URL+'/api/user/'+deleteObject.id;
    const options = {
      method: 'Delete',
      headers: headers
    }

    const response = await fetch(endpoint, options)

    const result = await response.json()
    router.push("/management/users")
    setDeleteModal(!deleteModal)
    setLoader(false)
  }
  
  return (
    <Row>
      <Col lg="12">
        <Link href="/management/users/create">
          <a>
            <Button className="btn btn-sm" color="primary" style={{marginBottom: '1em'}}>
            <i className="bi bi-people-fill me-2"></i> Tambah User
            </Button>
          </a>
        </Link>
        <div className="clearfix"></div>
        <UserTable header={header} rows={data} deleteToggle={deleteToggle}/>
      </Col>
      <Delete></Delete>
    </Row>
  );
};

export default Tables;
