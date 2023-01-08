import { useRouter } from 'next/router'
import Link from 'next/link';
import React, { useState } from 'react';
import { Row, Col, Modal, ModalHeader, ModalBody, ModalFooter, Button } from "reactstrap";
import UserTable from "../src/components/data/GenericTable";
import ButtonSubmitLoader from '../src/components/data/ButtonSubmitLoader';
import { getCookie } from 'cookies-next';

export const getServerSideProps = async (ctx) => {
  const token = getCookie('auth-token', {req: ctx.req, res: ctx.res})
  const headers = {
    'Content-Type': 'application/json',
    'Authorization': 'Bearer '+token
  }

  const res = await fetch(process.env.NEXT_SERVER_API_URL+'/api/notification', {headers: headers});
  const data = await res.json();

  if (res.status == 401) {
    return {
      redirect: {
          permanent: false,
          destination: '/login?errorCode=EXPIRED_SESSION',
      },
    }
  }

  const tableData = {
    title: "Notifikasi",
    subtitle: "",
    head: ["Pesan", "Tanggal"],
    show: ["message", "created_at"],
    rows: []
    };
    tableData['rows'] = data;


  return {
    props: {tableData: tableData, headers: headers}
  }
}

const Tables = ({tableData, headers}) => {
  const router = useRouter()
  const [loader, setLoader] = useState(false);

  return (
    <Row>
      <Col lg="12">
        <div className="clearfix"></div>

        <UserTable  data={tableData}/>
      </Col>
    </Row>
  );
};

export default Tables;
