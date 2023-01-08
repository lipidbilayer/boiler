import LoginLayout from "../src/layouts/LoginLayout"
import { deleteCookie, setCookie } from 'cookies-next';
import { useRouter } from 'next/router'
import moment from 'moment';
import { useEffect } from "react";

export const getServerSideProps = async (ctx) => {
    deleteCookie('auth-token', {req: ctx.req, res: ctx.res})

    return {
        // redirect: {
        //   permanent: false,
        //   destination: "/",
        // },
        props:{},
      };
}

const Verify = ({token}) => {
  const router = useRouter()
  useEffect(() => {
      if (window) { 
        window.localStorage.setItem('auth-token', "")
        router.push("/login")
      }
  }, []);
  return (<div></div>);
} 
  export default Verify