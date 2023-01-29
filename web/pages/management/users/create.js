import { useState } from 'react';
import { useRouter } from 'next/router'
import PageForm from "../../../src/components/data/PageForm";
import { getCookie } from 'cookies-next';

export const getServerSideProps = async (ctx) => {
  const header = {
    title: "Tambah User",
    icon: "bi bi-people me-2",
    fields: [
      {
        label: "Name",
        type: "text",
        name: "name",
        id: "name",
        required: 'true',
      },
      {
        label: "Username",
        type: "text",
        name: "username",
        id: "username",
        required: 'true',
      },
      {
        label: "Password",
        type: "password",
        name: "password",
        id: "password",
        required: 'true',
      }
    ],
    data: {},
  };

  return {
    props: {header: header}
  }
}



const UserDetail = ({header}) => {
  const router = useRouter()
  const [data, setData] = useState({});

  const handleSubmit = async (setLoader, event) => {
    event.preventDefault()
  }

  const changeData = (user) => {
    setData(user)
  }

  return (
      <PageForm inputData={header} handleSubmit={handleSubmit} changeDataCallback={changeData}></PageForm>
  )
}

export default UserDetail
