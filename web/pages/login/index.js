import LoginLayout from "../../src/layouts/LoginLayout"
import {
    Row,
    Col,
    Input,
    Label,
    Button,
    Alert,
  } from 'reactstrap';
import ButtonSubmitLoader from "../../src/components/data/ButtonSubmitLoader";
import { useState } from "react";
import { useRouter } from "next/router";

export const getServerSideProps = async (ctx) => {
    const errorCode = ctx.query.errorCode

    let errorMessage = null
    if(errorCode=="EXPIRED_SESSION"){
        errorMessage = "Sesi habis, Silahkan sign in"
    }
    return {
        props: {errorMessage: errorMessage}
      }
}

const Login = ({errorMessage}) => {
    const [loader, setLoader] = useState(false)
    const [request, setRequest] = useState({username: "", password: ""})
    const [alert, setAlert] = useState(errorMessage)
    const router = useRouter()

    const handleChange = e => {
        request[e.target.name] = e.target.value
    }

    const handleSubmit =  setLoader => async (event) => {
        event.preventDefault()
        const JSONdata = JSON.stringify(request)
        const endpoint = process.env.NEXT_PUBLIC_SERVER_API_URL+'/auth/login/';
        const options = {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSONdata,
        }
    
        const response = await fetch(endpoint, options)
    
        const result = await response.json()
        if(response.status == 401){
            setAlert("Login salah")
            return
        }

        router.push("login/verify?token="+result.token)
      }

    return (
    <Row>
        <Col className="col-sm-9 col-md-7 col-lg-5 mx-auto">
            <div className="card border-0 shadow rounded-3 my-5">
            <div className="card-body p-4 p-sm-5">
                <h3 className="card-title text-center mb-5 fw-light">Sign In</h3>
                <form>
                    <Alert color="danger" isOpen={alert != null} >
                        <i className="bi bi-x-circle-fill"> </i> {alert}
                    </Alert>
                    <div className="form-floating mb-3">
                        <Input type="text" className="form-control" id="username" name="username"  required={true} onChange={handleChange}/>
                        <Label for="username">Username</Label>
                    </div>
                    <div className="form-floating mb-3">
                        <Input type="password" className="form-control" id="password" name="password" required={true} onChange={handleChange} />
                        <Label for="password">Password</Label>
                    </div>

                    <div className="d-grid">
                        <Button className="btn btn-primary btn-login text-uppercase fw-bold" color='primary' disabled={loader} onClick={handleSubmit(setLoader)}>
                            <ButtonSubmitLoader text="Sign In" loader={loader}  />
                        </Button>
                    </div>
                </form>
            </div>
            </div>
        </Col>
    </Row>
    );
}

Login.getLayout = function getLayout(page) {
    return (
        <LoginLayout>{page}</LoginLayout>
    )
  }
  
  export default Login