import { Input, Container } from "reactstrap";
import backgroundImage from "../assets/images/background/background.jpg";


const LoginLayout = ({children}) => {
    
    return (
        <main>
            <div className="pageWrapper d-md-block d-lg-flex" style={{backgroundImage: `url(${backgroundImage.src})`, backgroundRepeat: 'no-repeat', backgroundSize: 'cover', backgroundPosition: 'center center'}}>
                <div className="contentArea">
                    <Container className="p-4 wrapper" fluid>
                        <div>{children}</div>
                    </Container>
                </div>
            </div>
        </main>
    );
}

export default LoginLayout;