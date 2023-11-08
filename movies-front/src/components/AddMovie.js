import React, { Fragment, useEffect, useState } from "react"

import { Button, Modal, Label, TextInput, Textarea, Select } from 'flowbite-react';
import { HiPlusCircle } from 'react-icons/hi';
import apiServer from "../consts"


const AddMovie = ({ onHide }) => {
    const [openModal, setOpenModal] = useState([]);
    const [directors, setDirectors] = useState([])

    const [title, setTitle] = useState([])
    const [description, setDescription] = useState([])
    const [year, setYear] = useState([])
    const [director_id, setDirectorId] = useState([])


    const props = { openModal, setOpenModal };

    const getDirectors = async () => {
        try {
            const response = await fetch(`${apiServer}/directors`)
            const jsonData = await response.json()

            setDirectors(jsonData)
        } catch (err) {
            console.error(err.message)
        }
    }

    const saveMovie = async e => {
        e.preventDefault()

        try {
            const body = { title, description, year, director_id }
            const response = await fetch(`${apiServer}/movies`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(body)
            })
            console.log("RESPONSE")
            console.log(response)

        } catch (err) {
            console.error(err.message)
        }
        props.setOpenModal(undefined)
        onHide()
    }

    useEffect(() => {
        getDirectors()
    }, [])

    const showModal = () => {
        getDirectors()
        props.setOpenModal('default')
    }


    return (
        <Fragment>
            <Button
                id="add-movie-btn"
                onClick={showModal}
                className="fixed z-90 bottom-10 right-8 rounded-full w-20 h-20">
                <HiPlusCircle className="w-16 h-16" />
            </Button>
            <Modal show={props.openModal === 'default'} onClose={() => props.setOpenModal(undefined)} >
                <Modal.Header>Add a Movie</Modal.Header>
                <Modal.Body>
                    <Label htmlFor="title" value="Title" />
                    <TextInput id="title" sizing="md" type="text" onChange={e => setTitle(e.target.value)} />
                    <Label htmlFor="description" value="Description" />
                    <Textarea id="description" onChange={e => setDescription(e.target.value)} />
                    <Label htmlFor="year" value="Year" />
                    <TextInput id="year" sizing="md" type="number" onChange={e => setYear(e.target.value)} />
                    <Label htmlFor="director_id" value="Director" />
                    <Select id="director_id" sizing="md" onChange={e => setDirectorId(e.target.value)} >
                        <option value={null} key={0}>
                            -- select director --
                        </option>
                        {directors.map(director =>
                            <option value={director.id} key={director.id}>
                                {director.name}
                            </option>
                        )}
                    </Select>
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={e => saveMovie(e)}>Add Movie</Button>
                </Modal.Footer>
            </Modal>
        </Fragment>
    )
}

export default AddMovie