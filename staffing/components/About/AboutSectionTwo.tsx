import Image from "next/image";

const AboutSectionTwo = () => {
  return (
    <section className="py-16 md:py-20 lg:py-28">
      <div className="container">
        <div className="-mx-4 flex flex-wrap items-center">
          <div className="w-full px-4 lg:w-1/2">
            <div
              className="wow fadeInUp relative mx-auto mb-12 aspect-[25/24] max-w-[500px] text-center lg:m-0"
              data-wow-delay=".15s"
            >
              <Image
                src="/images/about/about-image-2.svg"
                alt="about image"
                fill
              />
            </div>
          </div>
          <div className="w-full px-4 lg:w-1/2">
            <div className="wow fadeInUp max-w-[470px]" data-wow-delay=".2s">
              <div className="mb-9">
                <h3 className="mb-4 text-xl font-bold text-black dark:text-white sm:text-2xl lg:text-xl xl:text-2xl">
                  Do You Offer Direct Placement?
                </h3>
                <p className="text-base font-medium leading-relaxed text-body-color sm:text-lg sm:leading-relaxed">
                  Yes, we do! Contact us to discuss the positions you need to fill and a fair service rate.
                  All direct placement services are contingent upon hire.
                  We only have a one time engagement fee of $20 to confirm interest.
                  {/*Yes, we do! If you're interested in direct, contingent upon hire
                  placement solutions, please go&nbsp;
                  <a
                    href="https://Recruiting.ChristianHering.com/contact"
                    className="flex text-base text-dark group-hover:opacity-70 dark:text-white lg:mr-0 lg:inline-flex"
                  >
                    here
                  </a>
                  &nbsp;to contact our recruiting devision.*/}
                </p>
              </div>
              <div className="mb-9">
                <h3 className="mb-4 text-xl font-bold text-black dark:text-white sm:text-2xl lg:text-xl xl:text-2xl">
                  What Positions Do You Staff?
                </h3>
                <p className="text-base font-medium leading-relaxed text-body-color sm:text-lg sm:leading-relaxed">
                  We specialize in IT staffing but we also staff virtual assistants, sales specialists, managers, etc.
                  Reach out to see if we have candidates for the positions you're looking to fill today!
                </p>
              </div>
              <div className="mb-1">
                <h3 className="mb-4 text-xl font-bold text-black dark:text-white sm:text-2xl lg:text-xl xl:text-2xl">
                  Hidden Costs Of Direct Hire
                </h3>
                <p className="text-base font-medium leading-relaxed text-body-color sm:text-lg sm:leading-relaxed">
                  Directly hiring employees doesn't just cost their annual salary.
                  In fact, between your employee's salary, benefits, bonuses, and taxes, your employee will cost you&nbsp;
                  <a
                    href="https://web.mit.edu/e-club/hadzima/how-much-does-an-employee-cost.html"
                    className="flex text-base text-dark group-hover:opacity-70 dark:text-white lg:mr-0 lg:inline-flex"
                  >
                    1.25 to 1.4
                  </a>
                  &nbsp;times their base salary!
                  When you add in other hidden costs like internal recruitment, HR, job board fees, and more
                  complex accounting you'll start to see why direct placements aren't always the most economical option.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default AboutSectionTwo;
