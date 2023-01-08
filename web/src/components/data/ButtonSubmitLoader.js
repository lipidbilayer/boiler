
import { Spinner } from "reactstrap";

const ButtonSubmitLoader = ({text, loader}) => {
    // console.log(handleSubmit)
    if (loader){
      return (<span><Spinner
        as="span"
        animation="grow"
        size="sm"
        role="status"
        aria-hidden="true"
        /> Loading</span>);
    }
    return (<span>{text}</span>);
  }

  export default ButtonSubmitLoader;