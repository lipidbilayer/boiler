import { Button, Modal, ModalBody, ModalFooter, ModalHeader } from "reactstrap";
import ButtonSubmitLoader from "../data/ButtonSubmitLoader";

const Delete = ({deleteText, onFuncDelete, deleteModal, deleteToggle, loader}) => {  
    return (      
        <Modal isOpen={deleteModal} toggle={deleteToggle}>
            <ModalHeader toggle={deleteToggle}>Hapus User?</ModalHeader>
            <ModalBody>
            Hapus {deleteText}
            </ModalBody>
            <ModalFooter>
            <Button color="secondary" onClick={deleteToggle}>
                Cancel
            </Button>
            <Button color="danger" disabled={loader} onClick={onFuncDelete}>
                <ButtonSubmitLoader loader={loader} text="Delete" />
            </Button>
            </ModalFooter>
        </Modal>
    );
}

export default Delete;
