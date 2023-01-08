import Image from "next/image";
import parse from "html-react-parser";
import Link from "next/link";
import {Button, Card, CardBody, CardTitle, CardSubtitle, Table, A, Progress } from "reactstrap";
import moment, { utc } from 'moment';

const GenericTable = ({data, deleteToggle, detailToggle}) => {  

  const RenderAction = ({row, action, deleteToggle, detailToggle}) => {
    if(action.text == "Edit"){
      const url = action.url+row.id
      return (
      <Link key={action.text+row.id} href={url} passHref>
        <a>
          <Button className="btn btn-sm text-center" color={action.color}>
            <i className={action.icon}></i> {action.text}
          </Button>
        </a>
        </Link>
      )
    }else if(action.text == "Delete"){
      return (<Button key={"delete"+row.id} className="btn btn-sm text-center" color={action.color} onClick={(e) => deleteToggle({user: row})}>
        <i className={action.icon}></i> {action.text}
      </Button>)
    }else if(action.text == "Detail"){
      return (<Button key={"detail"+row.id} className="btn btn-sm text-center" color={action.color} onClick={(e) => detailToggle({data: row})}>
        <i className={action.icon}></i> {action.text}
      </Button>)
    }
    else{
      const url = action.url+row.id
      if(action.open_tab) {
        return (<a href={url} key={action.text+row.id}  className="btn btn-sm text-center btn-info" color={action.color} target="_blank"><i className={action.icon}></i> {action.text}</a>)
      }
      return (
        <Link key={action.text+row.id} href={url} passHref>
        <a>
          <Button className="btn btn-sm text-center" color={action.color}>
            <i className={action.icon}></i> {action.text}
          </Button>
        </a>
        </Link>
      )
    }
  }


  const RenderRow = ({type, row, deleteToggle}) => {  
    if(type == 'action')
       return (
         <td key={"action"+row.id}>
          <div className="button-group">
         {row[type].map((action) => {
           return RenderAction({row:row, action: action, deleteToggle: deleteToggle, detailToggle: detailToggle});
         })}
          </div>
          </td>
       );
    if(type == 'color_code'){
      return (
        <td style={{width: "5%"}} key={"action"+row.id}>
          <Progress
            value={0} style={{backgroundColor: row[type]}}
          />
        </td>
      )
    }

    if(type == "updated_at" || type == "created_at"){
      let date = utc(row[type], 'YYYY-MM-DDTHH:mm:ss[Z]').local().format('DD/MM/YYYY HH:mm');
      return (<td>{date}</td>)
    }
    
    let split_type = type.split(".")
    let nested_json = split_type.length > 1;
    if(nested_json){
      let nested_row = row;
      for(const st in split_type){
        if(nested_row) nested_row = nested_row[split_type[st]]
      }
      return (<td>{nested_row}</td>)
    }
  return (<td>{row[type]}</td>);
  }

  return (
    <Card>
      <CardBody>
        <CardTitle tag="h5">{data.title}</CardTitle>
        <CardSubtitle className="mb-2 text-muted" tag="h6">
          {data.subtitle}
        </CardSubtitle>
        <div className="table-responsive">
          <Table className="text-nowrap mt-3 align-middle" borderless>
            <thead>
              <tr>
              {data.head.map((tdata, index) => (
                <th key={tdata}>{tdata}</th>
              ))}
              </tr>
            </thead>
            <tbody>
              {data.rows.map((tdata, index) => (
                <tr key={index} className="border-top">
                      {data.show.map((key, index) => {
                        // ({key: key, row:tdata, deleteToggle}
                        return (<RenderRow key={index} type={key} row={tdata} deleteToggle={deleteToggle} />);
                })}
                </tr>
              ))}
            </tbody>
          </Table>
        </div>
      </CardBody>
    </Card>
  );
};

export default GenericTable;
