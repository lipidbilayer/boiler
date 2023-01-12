import LoginLayout from "../../src/layouts/LoginLayout"
import { setCookie } from 'cookies-next';
import { useRouter } from 'next/router'
import moment from 'moment';
import { useEffect } from "react";

export const getServerSideProps = async (ctx) => {
    const token = ctx.query.token
    let endOfDay = new Date();
    endOfDay.setHours(23,59,59,999);
    setCookie('auth_token', token, {req: ctx.req, res: ctx.res, expires: endOfDay, httpOnly: true})

    return {
        redirect: {
          permanent: false,
          destination: "/",
        },
      };
}

const Verify = ({data, token}) => {
  return (<div></div>);
} 
export default Verify