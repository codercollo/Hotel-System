<script setup lang="ts">
useHead({ title: "Contact Us — Where Luxury Meets Excellence" });

const api = useApi();
const form = reactive({
  name: "",
  email: "",
  phone: "",
  subject: "",
  message: "",
});
const submitted = ref(false);
const submitting = ref(false);
const submitError = ref<string | null>(null);

async function handleSubmit() {
  if (!form.name || !form.email || !form.subject || !form.message) {
    submitError.value = "Please fill in all required fields.";
    return;
  }
  submitting.value = true;
  submitError.value = null;
  try {
    await api.post("/api/v1/contact", { ...form });
    submitted.value = true;
  } catch (e: any) {
    submitError.value =
      e.message ?? "Failed to send message. Please try again.";
  } finally {
    submitting.value = false;
  }
}

const contactInfo = [
  {
    icon: "lucide:map-pin",
    label: "Address",
    value: "123 , South B, Nairobi 07102",
  },
  { icon: "lucide:phone", label: "Phone", value: "+1 (973) 555-0182" },
  { icon: "lucide:mail", label: "Email", value: "hello@luxuryhotel.com" },
  {
    icon: "lucide:clock",
    label: "Front Desk",
    value: "Open 24 hours, 7 days a week",
  },
];

const subjects = [
  "General Enquiry",
  "Room Reservation",
  "Special Occasions",
  "Spa & Wellness",
  "Events & Conferences",
  "Feedback",
];
</script>

<template>
  <div>
    <!-- Hero -->
    <section
      class="relative min-h-[50vh] flex items-end pb-16 bg-brand-900 overflow-hidden"
    >
      <div
        class="absolute inset-0 opacity-15"
        style="
          background-image: repeating-linear-gradient(
            45deg,
            white 0px,
            white 1px,
            transparent 1px,
            transparent 50%
          );
          background-size: 30px 30px;
        "
      />
      <div
        class="absolute inset-0"
        style="
          background: linear-gradient(
            135deg,
            rgb(25, 55, 48) 0%,
            rgb(38, 79, 72) 50%,
            rgb(52, 58, 32) 100%
          );
        "
      />
      <div class="relative z-10 max-w-7xl mx-auto section-px w-full pt-32">
        <div class="flex items-center gap-2 mb-4">
          <Icon
            v-for="i in 5"
            :key="i"
            name="lucide:star"
            class="w-3.5 h-3.5 text-accent fill-accent"
          />
        </div>
        <p class="eyebrow text-accent mb-3">Get in Touch</p>
        <h1
          class="font-display text-display-xl text-white font-light leading-tight max-w-xl"
        >
          We'd Love to<br />Hear From You
        </h1>
      </div>
    </section>

    <!-- Contact grid -->
    <section class="section-py section-px max-w-7xl mx-auto">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
        <!-- Left: info + map -->
        <div class="lg:col-span-1 space-y-8">
          <div>
            <p class="eyebrow mb-3">Contact Information</p>
            <h2
              class="font-display text-display-md text-brand-900 font-light mb-6"
            >
              Reach Us Anytime
            </h2>
            <ul class="space-y-5">
              <li
                v-for="info in contactInfo"
                :key="info.label"
                class="flex items-start gap-4"
              >
                <div
                  class="w-10 h-10 rounded-xl bg-forest/8 shrink-0 flex items-center justify-center mt-0.5"
                >
                  <Icon :name="info.icon" class="w-4 h-4 text-forest" />
                </div>
                <div>
                  <div
                    class="text-xs font-sans uppercase tracking-widest text-muted mb-0.5"
                  >
                    {{ info.label }}
                  </div>
                  <div class="font-sans text-sm text-brand-900">
                    {{ info.value }}
                  </div>
                </div>
              </li>
            </ul>
          </div>
          <!-- Map  -->
          <ClientOnly>
            <div class="rounded-2xl overflow-hidden aspect-[4/3]">
              <UiMap
                :lat="40.7357"
                :lng="-74.1724"
                label="Luxury Hotel — 123 Luxury Lane, Newark"
              />
            </div>
          </ClientOnly>

          <!-- Social -->
          <div>
            <p
              class="text-xs font-sans uppercase tracking-widest text-muted mb-3"
            >
              Follow Us
            </p>
            <div class="flex gap-3">
              <!-- Restore the opening <a> tag here -->
              <a
                v-for="social in [
                  'lucide:instagram',
                  'lucide:facebook',
                  'lucide:twitter',
                  'lucide:linkedin',
                ]"
                :key="social"
                href="#"
                class="w-9 h-9 rounded-xl border border-brand-200 flex items-center justify-center hover:border-accent hover:text-accent text-muted transition-colors duration-200"
              >
                <Icon :name="social" class="w-4 h-4" />
              </a>
            </div>
          </div>
        </div>

        <!-- Right: form -->
        <div class="lg:col-span-2">
          <!-- Success -->
          <div
            v-if="submitted"
            class="card p-10 flex flex-col items-center justify-center text-center min-h-[500px]"
          >
            <div
              class="w-16 h-16 rounded-full bg-forest/10 flex items-center justify-center mx-auto mb-6"
            >
              <Icon name="lucide:check" class="w-7 h-7 text-forest" />
            </div>
            <h3 class="font-display text-2xl text-brand-900 mb-3">
              Message Sent!
            </h3>
            <p
              class="text-sm text-muted font-sans max-w-sm leading-relaxed mb-6"
            >
              Thank you for reaching out. A member of our team will be in touch
              with you within 24 hours.
            </p>
            <UiButton variant="outline" size="sm" @click="submitted = false"
              >Send Another</UiButton
            >
          </div>

          <!-- Form -->
          <div v-else class="card p-8">
            <p class="eyebrow mb-2">Send a Message</p>
            <h2
              class="font-display text-display-sm text-brand-900 font-light mb-6"
            >
              How Can We Help?
            </h2>

            <div class="space-y-5">
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                <div>
                  <label
                    class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
                    >Full Name</label
                  >
                  <input
                    v-model="form.name"
                    type="text"
                    placeholder="Jenny Alexander"
                    class="input"
                  />
                </div>
                <div>
                  <label
                    class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
                    >Email Address</label
                  >
                  <input
                    v-model="form.email"
                    type="email"
                    placeholder="you@example.com"
                    class="input"
                  />
                </div>
              </div>

              <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                <div>
                  <label
                    class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
                    >Phone (optional)</label
                  >
                  <input
                    v-model="form.phone"
                    type="tel"
                    placeholder="+1 (555) 000-0000"
                    class="input"
                  />
                </div>
                <div>
                  <label
                    class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
                    >Subject</label
                  >
                  <select v-model="form.subject" class="input appearance-none">
                    <option value="" disabled>Select a subject</option>
                    <option v-for="s in subjects" :key="s" :value="s">
                      {{ s }}
                    </option>
                  </select>
                </div>
              </div>

              <div>
                <label
                  class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
                  >Message</label
                >
                <textarea
                  v-model="form.message"
                  rows="6"
                  placeholder="Tell us how we can make your stay perfect…"
                  class="input resize-none"
                />
              </div>

              <UiAlert
                v-if="submitError"
                variant="error"
                :message="submitError"
              />

              <UiButton
                size="md"
                class="w-full justify-center"
                :loading="submitting"
                :disabled="submitting"
                @click="handleSubmit"
              >
                Send Message
              </UiButton>

              <p class="text-xs text-muted font-sans text-center">
                We typically respond within 24 hours. For urgent requests,
                please call us directly.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- FAQ -->
    <section class="section-py section-px bg-surface-100">
      <div class="max-w-3xl mx-auto">
        <div class="text-center mb-10">
          <div class="ornament mb-3">✦</div>
          <p class="eyebrow mb-3">Quick Answers</p>
          <h2 class="font-display text-display-md text-brand-900 font-light">
            Frequently Asked
          </h2>
        </div>
        <div class="space-y-4">
          <details
            v-for="faq in [
              {
                q: 'What are your check-in and check-out times?',
                a: 'Check-in is from 3:00 PM and check-out is until 11:00 AM. Early check-in and late check-out can be arranged subject to availability.',
              },
              {
                q: 'Is parking available?',
                a: 'Yes, we offer complimentary valet parking for all hotel guests throughout their stay.',
              },
              {
                q: 'Do you offer airport transfers?',
                a: 'We provide complimentary luxury airport transfers. Please contact us at least 24 hours in advance to arrange your pickup.',
              },
              {
                q: 'Can I make a special occasion request?',
                a: 'Absolutely. Our concierge team can arrange flowers, champagne, bespoke room décor, and more. Simply mention it in your message.',
              },
            ]"
            :key="faq.q"
            class="card px-6 py-5 group"
          >
            <summary
              class="flex items-center justify-between cursor-pointer list-none"
            >
              <span class="font-display text-lg text-brand-900">{{
                faq.q
              }}</span>
              <Icon
                name="lucide:chevron-down"
                class="w-4 h-4 text-muted shrink-0 ml-4 group-open:rotate-180 transition-transform duration-200"
              />
            </summary>
            <p class="text-sm text-muted font-sans leading-relaxed mt-3 pb-1">
              {{ faq.a }}
            </p>
          </details>
        </div>
      </div>
    </section>
  </div>
</template>
