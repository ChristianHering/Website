"use client";
import { useState } from "react";
import SectionTitle from "../Common/SectionTitle";
import OfferList from "./OfferList";
import PricingBox from "./PricingBox";

const Pricing = () => {
  const [isHourly, setIsHourly] = useState(true);

  return (
    <section id="pricing" className="relative z-10 py-16 md:py-20 lg:py-28">
      <div className="container">
        <SectionTitle
          title="Simple and Affordable Pricing"
          paragraph="Take a look at some of our amazing bill rates for common positions!
          The final bill rate is dependant on the number of positions being filled and candidate requirements among others.
          Contact us to discuss your requirements and to get a quote for a position's hourly bill rate."
          center
          width="665px"
        />

        <div className="w-full">
          <div
            className="wow fadeInUp mb-8 flex justify-center md:mb-12 lg:mb-16"
            data-wow-delay=".1s"
          >
            <span
              onClick={() => setIsHourly(true)}
              className={`${
                isHourly
                  ? "pointer-events-none text-primary"
                  : "text-dark dark:text-white"
              } mr-4 cursor-pointer text-base font-semibold`}
            >
              Hourly
            </span>
            <div
              onClick={() => setIsHourly(!isHourly)}
              className="flex cursor-pointer items-center"
            >
              <div className="relative">
                <div className="h-5 w-14 rounded-full bg-[#1D2144] shadow-inner"></div>
                <div
                  className={`${
                    isHourly ? "" : "translate-x-full"
                  } shadow-switch-1 absolute left-0 top-[-4px] flex h-7 w-7 items-center justify-center rounded-full bg-primary transition`}
                >
                  <span className="active h-4 w-4 rounded-full bg-white"></span>
                </div>
              </div>
            </div>
            <span
              onClick={() => setIsHourly(false)}
              className={`${
                isHourly
                  ? "text-dark dark:text-white"
                  : "pointer-events-none text-primary"
              } ml-4 cursor-pointer text-base font-semibold`}
            >
              Yearly
            </span>
          </div>
        </div>

        <div className="grid grid-cols-1 gap-x-8 gap-y-10 md:grid-cols-2 lg:grid-cols-3">
          <PricingBox
            packageName="Customer Service"
            price={isHourly ? "20" : "41k"}
            duration={isHourly ? "hr" : "yr"}
            subtitle="Our friendly customer service representitives give your clients the support they deserve."
          >
            <OfferList text="Solves Non-Technical Problems" status="active" />
            <OfferList text="Handles Large Call Volumes" status="active" />
            <OfferList text="Email/Ticket Based Support" status="active" />
            <OfferList text="Can Follow A Sales Script" status="active" />
            <OfferList text="Available In Large Volume" status="active" />
            <OfferList text="May Lack A Batchelor's Degree" status="inactive" />
          </PricingBox>
          <PricingBox
            packageName="Virtual Assistant"
            price={isHourly ? "25" : "52k"}
            duration={isHourly ? "hr" : "yr"}
            subtitle="VA's lessen your workload and are a flexible option that can handle a variety of tasks."
          >
          <OfferList text="Email Management" status="active" />
          <OfferList text="Travel Planning" status="active" />
          <OfferList text="Personal Receptionist/Appointment Setter" status="active" />
          <OfferList text="Excel/Office Tools" status="active" />
          <OfferList text="Data Entry and Transcription" status="active" />
          <OfferList text="Can Work As A Customer Service Rep" status="active" />
          </PricingBox>
          <PricingBox
            packageName="Helpdesk/IT"
            price={isHourly ? "30" : "62k"}
            duration={isHourly ? "hr" : "yr"}
            subtitle="Our specialists give you the best customer service while fixing your issue the first time."
          >
            <OfferList text="Solid Technical Skills" status="active" />
            <OfferList text="Excellent Customer Service" status="active" />
            <OfferList text="Technical Writing/Documentation" status="active" />
            <OfferList text="Can Work Independently" status="active" />
            <OfferList text="VPN and Ticketing Experience" status="active" />
            <OfferList text="Can Work As A Virtual Assistant" status="active" />
          </PricingBox>
        </div>
      </div>

      <div className="absolute left-0 bottom-0 z-[-1]">
        <svg
          width="239"
          height="601"
          viewBox="0 0 239 601"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <rect
            opacity="0.3"
            x="-184.451"
            y="600.973"
            width="196"
            height="541.607"
            rx="2"
            transform="rotate(-128.7 -184.451 600.973)"
            fill="url(#paint0_linear_93:235)"
          />
          <rect
            opacity="0.3"
            x="-188.201"
            y="385.272"
            width="59.7544"
            height="541.607"
            rx="2"
            transform="rotate(-128.7 -188.201 385.272)"
            fill="url(#paint1_linear_93:235)"
          />
          <defs>
            <linearGradient
              id="paint0_linear_93:235"
              x1="-90.1184"
              y1="420.414"
              x2="-90.1184"
              y2="1131.65"
              gradientUnits="userSpaceOnUse"
            >
              <stop stopColor="#4A6CF7" />
              <stop offset="1" stopColor="#4A6CF7" stopOpacity="0" />
            </linearGradient>
            <linearGradient
              id="paint1_linear_93:235"
              x1="-159.441"
              y1="204.714"
              x2="-159.441"
              y2="915.952"
              gradientUnits="userSpaceOnUse"
            >
              <stop stopColor="#4A6CF7" />
              <stop offset="1" stopColor="#4A6CF7" stopOpacity="0" />
            </linearGradient>
          </defs>
        </svg>
      </div>
    </section>
  );
};

export default Pricing;
