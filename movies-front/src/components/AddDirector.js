import React, { Fragment, useEffect, useState } from "react"

import { Button, Modal, Label, TextInput } from 'flowbite-react';
import { HiPlusCircle } from 'react-icons/hi';
import apiServer from "../consts"


const AddDirector = ({ onHide }) => {
    const [openModal, setOpenModal] = useState([]);

    const [name, setName] = useState([])

    const props = { openModal, setOpenModal };

    useEffect(() => {
    }, [])

    const saveDirector = async e => {
        e.preventDefault()

        try {
            const body = { name }
            const response = await fetch(`${apiServer}/directors`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(body)
            })
            console.log(response)
        } catch (err) {
            console.error(err.message)
        }

        props.setOpenModal(undefined)
        onHide()
    }


    return (
        <>
            <Button
                id="add-director-btn"
                onClick={() => props.setOpenModal('default')}
                className="fixed z-90 bottom-10 right-8 rounded-full w-20 h-20">
                <HiPlusCircle className="w-16 h-16" />
            </Button>
            <Modal show={props.openModal === 'default'} onClose={() => props.setOpenModal(undefined)}>
                <Modal.Header>Add a Director</Modal.Header>
                <Modal.Body>
                    <Label htmlFor="name" value="Name" />
                    <TextInput id="name" sizing="md" type="text" onChange={e => setName(e.target.value)} autoComplete="off" />
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={e => saveDirector(e)}>Add Director</Button>
                </Modal.Footer>
            </Modal>
        </>
    )
}

export default AddDirector