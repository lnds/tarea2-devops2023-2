import '@testing-library/jest-dom'
import { fireEvent, render, screen } from '@testing-library/react';
import AddMovie from './AddMovie';

test('renders learn react link', () => {
    render(<AddMovie />);
    const linkElement = document.getElementById("add-movie-btn")
    expect(linkElement).toBeInTheDocument();
    fireEvent.click(linkElement)
    expect(screen.getByText("Add a Movie"))
});