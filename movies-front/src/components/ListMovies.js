import React, { Fragment, useEffect, useState } from "react"
import { Timeline } from 'flowbite-react';

import apiServer from "../consts";
import { TimelineItem } from "flowbite-react/lib/esm/components/Timeline/TimelineItem";
import AddMovie from "./AddMovie";

const ListMovies = () => {

    const [movies, setMovies] = useState([])

    const getMovies = async () => {
        try {
            const response = await fetch(`${apiServer}/movies`)
            const jsonData = await response.json()

            setMovies(jsonData)
        } catch (err) {
            console.error(err.message)
        }
    }


    const refreshMovies = async () => {
        console.log("refresh movies")
        await getMovies()
    }
    useEffect(() => {
        getMovies()
    }, [])

    return (
        <Fragment>
            <Timeline>
                {movies.map(movie =>
                    <TimelineItem key={movie.id}>
                        <Timeline.Point />
                        <Timeline.Content>
                            <Timeline.Time>
                                {movie.year}
                            </Timeline.Time>
                            <Timeline.Title>
                                {movie.title}
                            </Timeline.Title>
                            <Timeline.Body>
                                {movie.description}
                            </Timeline.Body>
                            <Timeline.Body>
                                <b>Director:</b> {movie.director.name}
                            </Timeline.Body>
                        </Timeline.Content>
                    </TimelineItem>
                )}
            </Timeline>
            <AddMovie onHide={refreshMovies} />
        </Fragment>
    )
}

export default ListMovies