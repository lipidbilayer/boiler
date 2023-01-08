import { useState } from 'react';
import { useRouter } from 'next/router'
import ButtonSubmitLoader from './ButtonSubmitLoader';
import InputColor from 'react-input-color';
import {
  Card,
  Row,
  Col,
  CardTitle,
  CardBody,
  Button,
  Form,
  FormGroup,
  Label,
  Input,
} from 'reactstrap';

const PageForm = ({inputData, handleSubmit, changeDataCallback}) => {
  const router = useRouter()
  const [loader, setLoader] = useState(false);
  const handleChange = e => {
    if(!e) return
    if(e.rgba){
      inputData.data.color_code = e.hex
      changeDataCallback(inputData.data)
      return
    }
    inputData.data[e.target.name] = e.target.value
    changeDataCallback(inputData.data)
  }

  const FieldGenerator = ({field}) => {
    switch(field.type) {
      case 'dropdown':
        return DropdownField(field, inputData.data, handleChange);
      case 'color_picker':
        return ColorPickerField(field, inputData.data, handleChange)
      default:
        return InputField(field, inputData.data, handleChange);
    }
  }

  const onClickSubmit = e => {
    setLoader(true)
    handleSubmit(setLoader, e)
  }

  return (
    <Row>
      <Col>
        <Button onClick={() => router.back()} className="btn btn-sm pull-right" color="warning" style={{marginBottom: '1em', float: 'right'}}>
          Kembali
        </Button>
        <div className="clearfix"></div>
        <Card>
          <CardTitle tag="h6" className="border-bottom p-3 mb-0">
            <i className={inputData.icon}> </i>
            {inputData.title}
          </CardTitle>
          <CardBody>
            <Form>
              {inputData.fields.map((field, index) => {
                  return (<FieldGenerator key={index} field={field} />);
              })}

              <FormGroup>
                <div className='modal-footer justify-content-center'>
                  <div className="button-group">
                    <Button  onClick={() => router.back()} >Batal</Button>
                    <Button type='submit' disabled={loader}  color='primary' onClick={onClickSubmit}>
                      <ButtonSubmitLoader text="Submit" loader={loader} />
                    </Button>
                  </div>
                </div>
              </FormGroup>

            </Form>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
};

const InputField = (field, data, onChangeFunc) => {
  const required = field.hasOwnProperty('required') ? 'required' : ''
  return (<FormGroup>
    { field.type != 'hidden' && <Label for={field.id}>{field.label}</Label>}
    <Input
      id={field.id}
      name={field.name}
      placeholder={field.placeholder}
      type={field.type} required={required} onChange={onChangeFunc} defaultValue={data[field.name]}
    />
  </FormGroup>)
}

const ColorPickerField = (field, data, onChangeFunc) => {
  const required = field.hasOwnProperty('required') ? 'required' : ''
  const colorHex = data[field.name] ? data[field.name] : '#000'
  return (
    <FormGroup>
      <Label for={field.id}>{field.label}</Label>
      <div>
      <InputColor
        id={field}
        name={field.name}
        initialValue={colorHex}
        onChange={onChangeFunc}
        placement="right"
        required={required}
      />
      </div>
    </FormGroup>
  )
}

const DropdownField = (field, data, onChangeFunc) => {
  return (
    <FormGroup>
    <Label for={field.id}>{field.label}</Label>
    <Input id={field.id} name={field.name} onChange={onChangeFunc} type="select"  defaultValue={data[field.name]}>
        {field.options.map((option) => {
            return (
              <option key={option.value} value={option.value} >{option.text}</option>
            )
        })}
    </Input>
  </FormGroup>
  )
}

export default PageForm;
