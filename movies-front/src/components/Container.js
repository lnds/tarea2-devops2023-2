import React, { Fragment } from "react"

const Container = ({ title, children }) => {
    return (
        <Fragment>
            <section className="bg-white dark:bg-gray-900">
                <div className="py-8 px-4 mx-auto max-w-screen-xl lg:px-12 sm:text-center lg:py-16">
                    <h2 className="mb-4 text-4xl tracking-tight font-extrabold text-gray-900 dark:text-white">
                        {title}
                    </h2>
                    <div className="pt-8 text-left sm:content-center">
                        {children}
                    </div>
                </div>
            </section>
        </Fragment>
    )
}

export default Container