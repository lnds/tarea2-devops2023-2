import '@testing-library/jest-dom'
import { fireEvent, render, screen } from '@testing-library/react';
import AddDirector from './AddDirector';

test('renders learn react link', () => {
    render(<AddDirector />);
    const linkElement = document.getElementById("add-director-btn")
    expect(linkElement).toBeInTheDocument();
    fireEvent.click(linkElement)
    expect(screen.getByText("Add a Movie"))
});