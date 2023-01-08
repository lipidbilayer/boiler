import LoginLayout from "../../src/layouts/LoginLayout"
import { setCookie } from 'cookies-next';
import { useRouter } from 'next/router'
import moment from 'moment';
import { useEffect } from "react";
import {UserProfileAPI} from "../../src/api/user/user";

export const getServerSideProps = async (ctx) => {
    const token = ctx.query.token
    let endOfDay = new Date();
    endOfDay.setHours(23,59,59,999);
    setCookie('auth-token', token, {req: ctx.req, res: ctx.res, expires: endOfDay, httpOnly: true})

    let data = await UserProfileAPI({token: token})

    return {
        // redirect: {
        //   permanent: false,
        //   destination: "/",
        // },
        props:{data: data, token: token},
      };
}

const Verify = ({data, token}) => {
  const router = useRouter()
  useEffect(() => {
      if (window) { 
        window.localStorage.setItem('auth-token', token)
        if(["admin", "owner"].includes(data.role)){
          router.push("/")
        }else{
          router.push("/scanner")
        }

      }
  }, []);
  return (<div></div>);
} 
  export default Verify