import React from 'react'
import { Dialog } from '@headlessui/react'

const Modal = ({ children, title, desc, isOpen, setIsOpen }: any) => {
    return (
        <Dialog open={isOpen} onClose={() => setIsOpen(false)} className="relative z-10">
            <div className="fixed inset-0 bg-black/30" aria-hidden="true" />
            <div className="fixed inset-0 flex items-center justify-center p-4">
                <Dialog.Panel className="w-[80%] rounded bg-white p-5">
                    <Dialog.Title className="text-3xl font-bold">{title}</Dialog.Title>
                    <Dialog.Description>
                        {desc}
                    </Dialog.Description>
                    <div>
                        {children}
                    </div>
                </Dialog.Panel>
            </div>
        </Dialog>
    )
}

export default Modal