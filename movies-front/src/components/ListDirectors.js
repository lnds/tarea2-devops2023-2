import React, { Fragment, useEffect, useState } from "react"
import { Timeline } from 'flowbite-react';

import apiServer from "../consts"
import { TimelineItem } from "flowbite-react/lib/esm/components/Timeline/TimelineItem";
import AddDirector from "./AddDirector";

const ListDirectors = () => {

    const [directors, setDirectors] = useState([])

    const getDirectors = async () => {
        try {
            const response = await fetch(`${apiServer}/directors`)
            const jsonData = await response.json()

            setDirectors(jsonData)
        } catch (err) {
            console.error(err.message)
        }
    }

    const refreshDirectors = async () => {
        console.log("refresh directors")
        await getDirectors()
    }

    useEffect(() => {
        getDirectors()
    }, [])

    return (
        <Fragment>
            <Timeline>
                {directors.map(director =>
                    <TimelineItem key={director.id}>
                        <Timeline.Point />
                        <Timeline.Content>
                            <Timeline.Title>
                                {director.name}
                            </Timeline.Title>
                        </Timeline.Content>
                    </TimelineItem>
                )}
            </Timeline>
            <AddDirector onHide={refreshDirectors} />
        </Fragment>
    )
}

export default ListDirectors